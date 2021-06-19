package holding_event

import (
	"context"
	"encoding/json"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type QuarterlyResultConsumer struct {
	ch      *amqp.Channel
	service QuarterlyResultService
}

type QuarterlyResultService interface {
	CalculateHolding(ctx context.Context, idAtivo int64) error
}

func NewQuarterlyResultConsumer(ch *amqp.Channel, service QuarterlyResultService) QuarterlyResultConsumer {
	return QuarterlyResultConsumer{
		ch: ch,
		service: service,
	}
}

func InitializeQuarterlyResultConsume(q QuarterlyResultConsumer, c *infrastructure.Cron) {
	log.Printf("Iniciando configuração de cron")
	err := c.Cron.AddFunc("0 0/1 * * * *", q.consume)
	if err != nil {
		log.Printf("Ocorreu um erro ao inicializar o consumer.")
	}
 }

func (q QuarterlyResultConsumer) consume() {
	log.Printf("Iniciando consumo de mensagens.")
	msgs, err := q.ch.Consume(
		infrastructure.GetResultQueueName(), 	    // queue
		"",                           // consumer
		false,                          // auto-ack
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

func (q QuarterlyResultConsumer) processMessage(ctx context.Context, body []byte) error {
	log.Printf("Mensage recebida: %s", body)
	result := &asset_domain.AssetQuarterlyResult{}
	err := json.Unmarshal(body, result)
	if err != nil {
		log.Printf("Erro ao converter evento")
		return err
	}
	err = q.service.CalculateHolding(ctx, result.Ativo)
	return err
}
