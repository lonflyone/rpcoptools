package mq

import (
	log "rpcoptool/pkg/logger"

	"github.com/streadway/amqp"
)

func SendMessage2MQ(body []byte, correlationID, callbackqueue string) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	log.LogrusObj.Infof("Operator 发送MQ队列:%v", callbackqueue)
	log.LogrusObj.Infof("Operator 发送确认序列:%v", correlationID)
	err = ch.Publish("", callbackqueue, false, false, amqp.Publishing{
		DeliveryMode:  amqp.Persistent,
		ContentType:   "application/json",
		CorrelationId: correlationID,
		Body:          body,
	})
	if err != nil {
		return
	}

	log.LogrusObj.Infoln("Operator 发送MQ成功...")
	return
}
