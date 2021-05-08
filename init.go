package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yyyThree/gin/config"
	"github.com/yyyThree/gin/library/rabbitmq"
	"github.com/yyyThree/gin/model/db"
	"runtime"
	"time"
)

// 初始化操作
func initService() {
	// 加载配置
	config.Load()
	fmt.Println("Port: ", config.Config.Http.Port)

	// 加载数据库配置
	db.Load()

	// 设置运行环境
	setEnv()

	// 设置CPU
	configRuntime()

	// 初始化rabbitmq
	initRabbitmq()
}

func setEnv() {
	if config.Config.App.Env == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func configRuntime() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	now := time.Now().String()
	fmt.Printf("Running time is %s\n", now)
	fmt.Printf("Running with %d CPUs\n", numCPU)
}

func initRabbitmq() {
	// 初始化MQ配置
	rabbitmq.InitConfig()

	// 声明队列
	rabbitmq.QueueDeclare()

	// 启动订阅
	rabbitmq.StartSubscriber()
}
