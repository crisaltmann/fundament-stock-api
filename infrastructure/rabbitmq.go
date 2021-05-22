package infrastructure

import (
	asset_repository "github.com/crisaltmann/fundament-stock-api/pkg/asset/repository"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func CreateRabbitMQCon() *amqp.Connection {
	amqUrl := os.Getenv("amqps://nstsjpmi:dN9SFZIn2R-MBDXeMaF63KKJbzB_0x6K@baboon.rmq.cloudamqp.com/nstsjpmi")
	if amqUrl == "" {
		panic("URL RabbitMQ nao encontrada.")
	}
	conn, err := amqp.Dial("amqps://xruxgkhh:7z2_613ze4F7qjbjfbkE43-hHuQ8_YaT@baboon.rmq.cloudamqp.com/xruxgkhh")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	return conn
}

func CreateRabbitMQChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()
	return ch
}

func ConfigureQueue(ch *amqp.Channel) {
	configureQuarterlyResultQueue(ch)
}

func configureQuarterlyResultQueue(ch *amqp.Channel) {
	_, err := ch.QueueDeclare(
		asset_repository.ResultQueueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func teste() {
	//amqps://xruxgkhh:7z2_613ze4F7qjbjfbkE43-hHuQ8_YaT@baboon.rmq.cloudamqp.com/xruxgkhh
	conn, err := amqp.Dial("amqps://xruxgkhh:7z2_613ze4F7qjbjfbkE43-hHuQ8_YaT@baboon.rmq.cloudamqp.com/xruxgkhh")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
