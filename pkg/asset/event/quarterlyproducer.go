package event

import (
	"encoding/json"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type QuarterlyResultProducer struct {
	conn		*amqp.Connection
}

func NewQuarterlyResultProducer(conn *amqp.Connection) QuarterlyResultProducer {
	return QuarterlyResultProducer{
		conn: conn,
	}
}

func (q QuarterlyResultProducer) PublishQuarterlyResultEvent(result asset_domain.AssetQuarterlyResult) error {
	ch, err := q.conn.Channel()
	infrastructure.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	body, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("Erro ao publicar mensagem na fila.", err)
	}
	err = ch.Publish(
		"",                              // exchange
		infrastructure.GetResultQueueName(),           // routing key
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