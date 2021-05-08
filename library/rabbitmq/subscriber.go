package rabbitmq

import "github.com/yyyThree/gin/library/rabbitmq/subscriber"

// 消息订阅
// 需要启动的的订阅者放这里
var Subscribers = []func(){
	// 商品模块
	subscriber.SyncItemSearch,

	// 死信队列
	subscriber.ItemDl,
}
