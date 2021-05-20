package main

import (
	"github.com/crisaltmann/fundament-stock-api/cmd/api"
	"github.com/crisaltmann/fundament-stock-api/config"
	"github.com/crisaltmann/fundament-stock-api/infrastructure"
	asset_sync2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/asset-sync"
	"github.com/crisaltmann/fundament-stock-api/server"
	"github.com/streadway/amqp"
	"go.uber.org/fx"
	"log"
)

// @title Fundament Stock Api Swagger API
// @version 1.0
// @description Swagger API for Fundament Stock Api.
// @termsOfService http://swagger.io/terms/

// @contact.name Cristiano Altmann
// @contact.email crisaltmann@gmail.com

// @license.name MIT
// @license.url https://github.com/crisaltmann/fundament-stock-api

// @BasePath /
func main() {
	log.Println("Iniciando...")


	//teste()

	app := fx.New(
		infrastructure.Module,
		config.Module,
		server.Module,
		api.Asset,
		api.Order,
		api.Portfolio,
		api.Quarter,
		asset_sync2.Module,
		fx.Invoke(
			server.InitServer,
		),
	)
	app.Run()
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
