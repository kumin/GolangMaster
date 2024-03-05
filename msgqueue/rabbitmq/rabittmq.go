package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMsg() {
	conn, err := amqp.Dial(RabbitMqHost)
	FailOnError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		HelloQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "hello world!"
	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	FailOnError(err)
	log.Printf("message '%s' was sent", body)
}

func ConsumeMsg() {
	conn, err := amqp.Dial(RabbitMqHost)
	FailOnError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		HelloQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err)
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	var forever chan struct{}
	go func() {
		for {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}
	}()

	<-forever
}
