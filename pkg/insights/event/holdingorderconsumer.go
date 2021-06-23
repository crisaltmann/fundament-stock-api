package insight_event

import (
	"context"
	"encoding/json"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type InsightConsumer struct {
	conn		*amqp.Connection
	service    InsightService
}

type InsightService interface {
	CalculateInsights(ctx context.Context, holdings holding_domain.Holdings) error
}

func NewInsightConsumer(conn *amqp.Connection, service InsightService) InsightConsumer {
	return InsightConsumer{
		conn: conn,
		service: service,
	}
}

func InitializeInsightConsume(q InsightConsumer, c *infrastructure.Cron) {
	log.Printf("Iniciando configuração de cron")
	err := c.Cron.AddFunc("0 0/1 * * * *", q.consume)
	if err != nil {
		log.Printf("Ocorreu um erro ao inicializar o consumer.")
	}
 }

func (q InsightConsumer) consume() {
	log.Printf("Iniciando consumo de mensagens.")

	ch, err := q.conn.Channel()
	infrastructure.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		infrastructure.GetHoldingResultQueueName(), 	       // queue
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
			ctx := context.TODO()
			err := q.processMessage(ctx, d.Body)
			if err != nil {
				log.Printf("Ocorreu um erro ao processar a mensagem. %v", err)
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

func (q InsightConsumer) processMessage(ctx context.Context, body []byte) error {
	log.Printf("Mensagem recebida: %s", body)
	holdings := holding_domain.Holdings{}
	err := json.Unmarshal(body, &holdings)
	if err != nil {
		log.Printf("Erro ao converter evento")
		return err
	}
	err = q.service.CalculateInsights(ctx, holdings)
	return err
}
