package infrastructure

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func CreateRabbitMQCon() *amqp.Connection {
	amqUrl := os.Getenv("CLOUDAMQP_URL")
	if amqUrl == "" {
		log.Printf("URL RabbitMQ nao encontrada. Usando valor default")
		amqUrl = "amqps://nstsjpmi:dN9SFZIn2R-MBDXeMaF63KKJbzB_0x6K@baboon.rmq.cloudamqp.com/nstsjpmi"
	}
	conn, err := amqp.Dial(amqUrl)
	FailOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	return conn
}

func CreateRabbitMQChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	//defer ch.Close()
	return ch
}

func ConfigureQueue(ch *amqp.Channel) {
	configureQuarterlyResultQueue(ch)
	configureOrderQueue(ch)
}

func configureQuarterlyResultQueue(ch *amqp.Channel) {
	_, err := ch.QueueDeclare(
		ResultQueueName, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	FailOnError(err, "Failed to declare a queue")
}

func configureOrderQueue(ch *amqp.Channel) {
	_, err := ch.QueueDeclare(
		OrderQueueName, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	FailOnError(err, "Failed to declare a queue")
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
