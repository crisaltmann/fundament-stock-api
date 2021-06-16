package holding_event

import (
	"encoding/json"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	order_domain "github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type HoldingOrderConsumer struct {
	conn		*amqp.Connection
	service    HoldingOrderService
}

type HoldingOrderService interface {
	CalculateHolding(idAtivo int64) error
}

func NewHoldingOrderConsumer(conn *amqp.Connection, service HoldingOrderService) HoldingOrderConsumer {
	return HoldingOrderConsumer{
		conn: conn,
		service: service,
	}
}

func InitializeOrderConsume(q HoldingOrderConsumer, c *infrastructure.Cron) {
	log.Printf("Iniciando configuração de cron")
	err := c.Cron.AddFunc("0 0/1 * * * *", q.consume)
	if err != nil {
		log.Printf("Ocorreu um erro ao inicializar o consumer.")
	}
 }

func (q HoldingOrderConsumer) consume() {
	log.Printf("Iniciando consumo de mensagens.")

	ch, err := q.conn.Channel()
	infrastructure.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		infrastructure.OrderQueueName, 	       // queue
		"",                           // consumer
		false,                         // auto-ack
		false,                        // exclusive
		false,                         // no-local
		false,                          // no-wait
		nil,                              // args
	)
	if err != nil {
		log.Printf("Erro ao consumir mensagens.")
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			err := q.processMessage(d.Body)
			if err != nil {
				log.Printf("Ocorreu um erro ao processar a mensagem.")
				//Adicionar mecanismo de tratativas
				d.Reject(true)
			} else {
				d.Ack(false)
			}
		}
	}()

	<-forever
	log.Printf("Encerrando ciclo de consumo de mensagens.")
}

func (q HoldingOrderConsumer) processMessage(body []byte) error {
	log.Printf("Mensage recebida: %s", body)
	result := &order_domain.Order{}
	err := json.Unmarshal(body, result)
	if err != nil {
		log.Printf("Erro ao converter evento")
		return err
	}
	err = q.service.CalculateHolding(result.Ativo)
	return err
}
