package operator

import (
	"context"
	"encoding/json"
	"rpcoptool/app/operator/repository/mq"
	log "rpcoptool/pkg/logger"
	"rpcoptool/proto/pb"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskCreate(ctx context.Context) error {
	msgs, err := mq.ReciveMessage(ctx)
	if err != nil {
		return err
	}
	var forever chan struct{}
	log.LogrusObj.Infof("Operator 开始接收")
	go func() {
		//block住，一直执行
		for d := range msgs {
			log.LogrusObj.Infof("Operator run Task: %s", d.Body)

			// 简单处理
			TaskModel := new(pb.TaskModel)
			err = json.Unmarshal(d.Body, TaskModel)
			if err != nil {
				log.LogrusObj.Errorf("Operator run Task: fail %s", err)
			}
			TaskModel.Content = "完成了RPC的Operator调用"

			// 发送callbackqueue
			body, _ := json.Marshal(TaskModel)
			err = mq.SendMessage2MQ(body, d.CorrelationId, d.ReplyTo)
			if err != nil {
				log.LogrusObj.Infof("发送回调消息失败: %s", err)
			}
			d.Ack(false)

		}
	}()

	log.LogrusObj.Infoln(err)
	//主进程block住,使得该线程一直能一直执行
	<-forever

	return nil
}
