package subscriber

import (
	"fmt"
	"gin/library/log"
	"gin/library/rabbitmq/common"
	"github.com/streadway/amqp"
	"github.com/yyyThree/rabbitmq"
	"github.com/yyyThree/zap"
)

// 同步商品搜索数据
func SyncItemSearch()  {
	queue := common.SyncItemSearch
	go func() {
		err := rabbitmq.Subscribe(queue.Name, func(msg amqp.Delivery) {
			fmt.Println("SyncItemSearch", string(msg.Body))
			rabbitmq.Reject(msg)
		})
		if err != nil {
			log.GetLogger().Info("SyncItemSearch", zap.BaseMap{
				"queue": queue,
				"error": err,
			})
			fmt.Println("SyncItemSearch err：", err.Error())
		}
	}()
}