package subscriber

import (
	"fmt"
	"gin/library/log"
	"gin/library/rabbitmq/common"
	"github.com/streadway/amqp"
	"github.com/yyyThree/rabbitmq"
	"github.com/yyyThree/zap"
)

// 通用死信队列订阅
func CommonDl()  {
	queue := common.DlQueue
	go func() {
		err := rabbitmq.Subscribe(queue.Name, func(msg amqp.Delivery) {
			fmt.Println("CommonDl", string(msg.Body))
			rabbitmq.Ack(msg)
		})
		if err != nil {
			log.GetLogger().Info("CommonDl", zap.BaseMap{
				"queue": queue,
				"error": err,
			})
			fmt.Println("CommonDl err：", err.Error())
		}
	}()
}