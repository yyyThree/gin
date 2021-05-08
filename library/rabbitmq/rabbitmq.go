package rabbitmq

import (
	"fmt"
	"github.com/yyyThree/gin/config"
	"github.com/yyyThree/gin/library/log"
	"github.com/yyyThree/gin/library/rabbitmq/common"
	"github.com/yyyThree/rabbitmq"
	"github.com/yyyThree/zap"
)

// 初始化rabbitmq配置
func InitConfig() {
	_ = rabbitmq.InitConfig(rabbitmq.Config{
		Base: rabbitmq.Base{
			Host:     config.Config.Rabbitmq.Host,
			Port:     config.Config.Rabbitmq.Port,
			User:     config.Config.Rabbitmq.User,
			Password: config.Config.Rabbitmq.Password,
			Vhost:    config.Config.Rabbitmq.Vhost,
		},
		Exchange: rabbitmq.Exchange{
			Direct:      config.Config.Rabbitmq.ExDirect,
			Topic:       config.Config.Rabbitmq.ExTopic,
			DeathLetter: config.Config.Rabbitmq.ExDeathLetter,
		},
		Ttl: rabbitmq.Ttl{
			QueueMsg: config.Config.Rabbitmq.TtlQueueMsg,
			Msg:      config.Config.Rabbitmq.TtlMsg,
		},
		Admin: rabbitmq.Admin{
			User:     config.Config.Rabbitmq.AdminUser,
			Password: config.Config.Rabbitmq.AdminPassword,
		},
		Log: rabbitmq.Log{
			Dir: config.Config.Rabbitmq.LogDir,
		},
	})
}

// 声明队列
func QueueDeclare() {
	// 带死信参数的队列声明
	for _, queue := range common.QueueDirectWithDlList {
		queue := queue
		err := rabbitmq.QueueDeclareWithDl(queue.Name, queue.Keys, queue.DlxKey)
		if err != nil {
			log.GetLogger().Info("QueueDirectWithDl", zap.BaseMap{
				"queue": queue,
				"error": err,
			})
			fmt.Println("QueueDirectWithDl err：", err.Error())
		}
	}

	// 死信队列声明
	for _, queue := range common.QueueDlList {
		queue := queue
		err := rabbitmq.QueueDeclareDl(queue.Name, queue.Keys)
		if err != nil {
			log.GetLogger().Info("QueueDeclareDl", zap.BaseMap{
				"queue": queue,
				"error": err,
			})
			fmt.Println("QueueDeclareDl err：", err.Error())
		}
	}
}

// 启动订阅
func StartSubscriber() {
	for _, subscriber := range Subscribers {
		subscriber()
	}
}
