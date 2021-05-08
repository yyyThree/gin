package subscriber

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yyyThree/gin/library/log"
	"github.com/yyyThree/gin/library/rabbitmq/common"
	"github.com/yyyThree/rabbitmq"
	"github.com/yyyThree/zap"
)

// 商品通用死信队列订阅
func ItemDl() {
	queue := common.ItemDlQueue
	go func() {
		err := rabbitmq.Subscribe(queue.Name, func(msg amqp.Delivery) {
			fmt.Println("ItemDl", string(msg.Body))
			rabbitmq.Ack(msg)
		})
		if err != nil {
			log.GetLogger().Info("ItemDl", zap.BaseMap{
				"queue": queue,
				"error": err,
			})
			fmt.Println("ItemDl err：", err.Error())
		}
	}()
}
