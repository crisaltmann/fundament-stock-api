package order_event

import (
	"encoding/json"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	order_domain "github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type OrderProducer struct {
	conn		*amqp.Connection
}

func NewOrderProducer(conn *amqp.Connection) OrderProducer {
	return OrderProducer{
		conn: conn,
	}
}

func (q OrderProducer) PublishOrderEvent(result order_domain.Order) error {
	ch, err := q.conn.Channel()
	infrastructure.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	body, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("Erro ao publicar evento de movimentação na fila.", err)
	}
	err = ch.Publish(
		"",                             // exchange
		infrastructure.GetOrderQueueName(),           // routing key
		false,                         // mandatory
		false,                         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("Erro ao publicar mensagem na fila.", err)
	}
	log.Print("Evento movimentacao enviado com sucesso.")
	return nil
}