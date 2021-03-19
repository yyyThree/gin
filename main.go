package main

import (
	"fmt"
	"gin/config"
	"gin/model/db"
	"gin/router"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 初始化操作
func init() {
	// 加载配置
	config.Load()
	// 加载数据库配置
	db.Load()
	// 设置运行环境
	if config.Config.App.Env == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// 启动服务器
func main() {
	// 启动http服务器
	r := router.Router()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Config.Http.Port),
		Handler:        r,
		ReadTimeout:    time.Duration(config.Config.Http.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Config.Http.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 使用manner做平滑重启（兼容window）
	err := manners.NewWithServer(s).ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("http服务器启动失败: %s \n", err))
	}
}
