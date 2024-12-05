package mq

import (
	"strings"

	"github.com/streadway/amqp"
)

var RabbitMq *amqp.Connection

func InitRabbitMQ() {
	connString := strings.Join([]string{"amqp", "://", "guest", ":", "guest", "@", "localhost", ":", "5672", "/"}, "")
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	RabbitMq = conn
}
