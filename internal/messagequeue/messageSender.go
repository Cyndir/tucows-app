package messagequeue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Cyndir/tucows-app/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

const QUEUE_NAME = "payment"

//go:generate mockgen -destination=../mocks/messageSender.go -package=mocks -source=messageSender.go
type MessageQueue interface {
	Publish(message model.Order) error
	Close()
}

type messageQueueImpl struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New() MessageQueue {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	_, err = ch.QueueDeclare(
		QUEUE_NAME, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &messageQueueImpl{
		conn: conn,
		ch:   ch,
	}
}

func (mq *messageQueueImpl) Publish(message model.Order) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = mq.ch.PublishWithContext(context.Background(),
		"",         // exchange
		QUEUE_NAME, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})

	log.Printf("Sent message")
	return err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (mq *messageQueueImpl) Close() {
	mq.conn.Close()
	mq.ch.Close()
}
