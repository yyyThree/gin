package common

type Queue struct {
	Name string // 队列名
	Keys []string // 队列绑定的路由键值
	DlxKey string // 队列绑定的死信队列路由
}

// 带死信参数的直连交换机队列
var QueueDirectWithDlList = []Queue{SyncItemSearch}

var (
	SyncItemSearch = Queue{
		Name:   "syncItemSearch",
		Keys:   []string{ItemAdd},
		DlxKey: CommonDl,
	}
)

// 死信队列
var QueueDlList = []Queue{DlQueue}

var (
	DlQueue = Queue{
		Name: "commonDl",
		Keys: []string{CommonDl},
	}
)