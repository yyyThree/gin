package rabbitmq

import "gin/library/rabbitmq/subscriber"

// 消息订阅
// 需要订阅的消息放这里
var Subscribers = []func(){
	// 商品模块
	subscriber.SyncItemSearch,

	// 死信队列
	subscriber.CommonDl,
}