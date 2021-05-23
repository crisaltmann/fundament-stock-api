package event

import (
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"github.com/streadway/amqp"
	"log"
)

type QuarterlyResultConsumer struct {
	ch		*amqp.Channel
}

func NewQuarterlyResultConsumer(ch *amqp.Channel) QuarterlyResultConsumer {
	return QuarterlyResultConsumer{
		ch: ch,
	}
}

func InitializeConsume(q QuarterlyResultConsumer, c *infrastructure.Cron) {
	//c := cron.New()
	err := c.Cron.AddFunc("0 0/1 * * * *", q.consume)
	if err != nil {
		log.Println("Ocorreu um erro ao inicializar o consumer.")
	}
	//c.Start()
 }

func (q QuarterlyResultConsumer) consume() {
	msgs, err := q.ch.Consume(
		infrastructure.ResultQueueName, // queue
		"",                             // consumer
		true,                           // auto-ack
		false,                          // exclusive
		false,                          // no-local
		false,                          // no-wait
		nil,                            // args
	)
	if err != nil {
		log.Println("Erro ao consumir mensagens.")
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever
}
