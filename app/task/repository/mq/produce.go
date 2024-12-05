package mq

import (
	"math/rand"
	log "rpcoptool/pkg/logger"
	"time"

	"github.com/streadway/amqp"
)

func randomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func SendMessage2MQ(body []byte) (correlationID, callbackQueue string, err error) {
	correlationID = randomString(32)
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	q, _ := ch.QueueDeclare("", true, false, false, false, nil)
	callbackQueue = q.Name
	err = ch.Publish("", "rpc_queue", false, false, amqp.Publishing{
		DeliveryMode:  amqp.Persistent,
		ContentType:   "application/json",
		ReplyTo:       callbackQueue,
		CorrelationId: correlationID,
		Body:          body,
	})
	if err != nil {
		return
	}

	log.LogrusObj.Infoln("Task 发送MQ成功...")
	ch.Close()
	return
}
