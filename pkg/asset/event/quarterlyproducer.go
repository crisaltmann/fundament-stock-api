package event

import (
	"encoding/json"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/streadway/amqp"
	"log"
)

type QuarterlyResultProducer struct {
	ch		*amqp.Channel
}

func NewQuarterlyResultProducer(ch *amqp.Channel) QuarterlyResultProducer {
	return QuarterlyResultProducer{
		ch: ch,
	}
}

func (q QuarterlyResultProducer) PublishQuarterlyResultEvent(result asset_domain.AssetQuarterlyResult) error {
	body, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("Erro ao publicar mensagem na fila.", err)
	}
	err = q.ch.Publish(
		"",                             // exchange
		infrastructure.ResultQueueName, // routing key
		false,                          // mandatory
		false,                          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("Erro ao publicar mensagem na fila.", err)
	}
	log.Print("Evento enviado com sucesso.")
	return nil
}