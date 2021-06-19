package holding_event

import (
	"encoding/json"
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	"github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type HoldingResultProducer struct {
	conn		*amqp.Connection
}

func NewHoldingResultProducer(conn *amqp.Connection) HoldingResultProducer {
	return HoldingResultProducer{
		conn: conn,
	}
}

func (q HoldingResultProducer) PublishHoldingResultEvent(event holding_domain.Holdings) error {
	removeAssetLogo(&event)
	ch, err := q.conn.Channel()
	infrastructure.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Erro ao publicar evento de resultados de holding na fila.", err)
	}
	err = ch.Publish(
		"",                             // exchange
		infrastructure.HoldingResultQueueName,           // routing key
		false,                         // mandatory
		false,                         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("Erro ao publicar mensagem na fila.", err)
	}
	log.Print("Evento resultado de holding enviado com sucesso.")
	return nil
}

func removeAssetLogo(event *holding_domain.Holdings) {
	for i := 0; i < len(event.Holdings); i++ {
		for j := 0; j < len(event.Holdings[i].HoldingsAtivo); j++ {
			holdingAtivo := event.Holdings[i].HoldingsAtivo[j]
			holdingAtivo.Ativo.Logo = ""
		}
	}
}