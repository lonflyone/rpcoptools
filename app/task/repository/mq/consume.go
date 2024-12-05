package mq

import (
	"context"
	"encoding/json"
	"rpcoptool/app/task/repository/db/dao"
	"rpcoptool/app/task/repository/db/model"
	log "rpcoptool/pkg/logger"
	"rpcoptool/proto/pb"

	"github.com/streadway/amqp"
)

func ReciveMessage(ctx context.Context, queueName string) (msgs <-chan amqp.Delivery, err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	ch.Qos(1, 0, false)
	msgs, err = ch.Consume(queueName, "", false, false, false, false, nil)

	return
}

func GetTaskOpMessage(ctx context.Context, Callbackqueue, correlationID string) error {
	msgs, err := ReciveMessage(ctx, Callbackqueue)
	if err != nil {
		return err
	}
	for d := range msgs {
		if d.CorrelationId == correlationID {
			log.LogrusObj.Infof("Task Received Operator Task: %v", Callbackqueue)
			log.LogrusObj.Infof("Task Received run Task: %s", d.Body)
			// 落库
			model := new(model.TaskOp)
			taskModel := new(pb.TaskModel)
			err = json.Unmarshal(d.Body, taskModel)
			if err != nil {
				log.LogrusObj.Infof("Task Received run Task Fail: %s", err)
			}
			model.CorrelationID = correlationID
			model.TaskId = uint(taskModel.Id)
			model.OpId = 1
			model.RpcQueue = "rpc_queue"
			model.Status = 1
			err = dao.NewTaskDao(ctx).CreateOpT(model)
			if err != nil {
				log.LogrusObj.Errorf("CreateOpT fail %v", err)
				return err
			}
			err = dao.NewTaskDao(ctx).UpdateTask(taskModel)
			if err != nil {
				log.LogrusObj.Errorf("Update fail %v", err)
				return err
			}
			d.Ack(false)
			break
		}
	}
	return nil
}
