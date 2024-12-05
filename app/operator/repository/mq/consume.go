package mq

import (
	"context"

	"github.com/streadway/amqp"
)

func ReciveMessage(ctx context.Context) (msgs <-chan amqp.Delivery, err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	q, _ := ch.QueueDeclare("rpc_queue", true, false, false, false, nil)
	ch.Qos(1, 0, false)
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
