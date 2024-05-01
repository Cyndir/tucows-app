//go:build integration

package messagequeue

import (
	"encoding/json"
	"testing"

	"github.com/Cyndir/tucows-app/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestMessageQueue(t *testing.T) {

	sender := New()
	consumer := buildConsumer(t)
	order := model.Order{
		ID:         "test",
		Total:      1000,
		CustomerID: "1",
		ProductID:  "1",
	}
	err := sender.Publish(order)
	assert.NoError(t, err)

	msgs, err := consumer.Consume(
		"payment", // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	assert.NoError(t, err)

	msg := <-msgs

	var actualOrder model.Order
	json.Unmarshal(msg.Body, &actualOrder)
	assert.Equal(t, order, actualOrder)
}

func buildConsumer(t *testing.T) *amqp.Channel {
	t.Helper()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	assert.NoError(t, err)
	ch, err := conn.Channel()
	assert.NoError(t, err)

	_, err = ch.QueueDeclare(
		"payment", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	assert.NoError(t, err)
	return ch
}
