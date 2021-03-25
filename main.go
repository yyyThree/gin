package main

import (
	"context"
	"fmt"
	"gin/config"
	"gin/model/db"
	"gin/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	go func() {
		// 服务连接
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http服务器启动失败: %s\n", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置等待时间
	quit := make(chan os.Signal) // 创建一个接收信号的通道
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit // 阻塞，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个最大等待时间的context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.Http.ShutdownTimeOut) * time.Second)
	defer cancel()
	// 优雅关闭服务（将未处理完的请求处理完再关闭服务），超过等待时间就超时退出
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
