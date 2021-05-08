package subscriber

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yyyThree/gin/helper"
	"github.com/yyyThree/gin/library/log"
	"github.com/yyyThree/gin/library/rabbitmq/common"
	"github.com/yyyThree/gin/model/param"
	"github.com/yyyThree/gin/service"
	"github.com/yyyThree/rabbitmq"
	"github.com/yyyThree/zap"
)

// 同步商品搜索数据
func SyncItemSearch() {
	queue := common.SyncItemSearch
	go func() {
		_ = rabbitmq.Subscribe(queue.Name, func(msg amqp.Delivery) {
			params := param.ItemSync{}
			_ = json.Unmarshal(msg.Body, &params)

			if helper.HasAnyEmpty(params.ItemId, params.SyncType) {
				rabbitmq.Reject(msg)
				return
			}
			err := (&service.ItemSearch{}).Sync(params)
			if err != nil {
				log.GetLogger().Info("SyncItemSearch", zap.BaseMap{
					"queue":  queue,
					"params": params,
					"error":  err,
				})
				fmt.Println("SyncItemSearch err：", err.Error())
				rabbitmq.Nack(msg)
				return
			}
			rabbitmq.Ack(msg)
		})
	}()
}
