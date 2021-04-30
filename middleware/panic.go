// 500错误统一监听处理
package middleware

import (
	"fmt"
	"gin/library/log"
	"gin/output"
	"gin/output/code"
	"github.com/gin-gonic/gin"
	"github.com/yyyThree/zap"
	"runtime"
)

type Panic struct {
}

// 统一错误处理
// TODO @2021-02-09 后续可使用 CustomRecovery 方法，待gin框架发布新版
func (panic *Panic) Handler() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				// TODO 邮件报警
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				fmt.Printf("==> %s\n", string(buf[:n]))
				log.GetLogger().Panic("panic", zap.BaseMap{
					"err": e,
				})
				output.Response(c, nil, output.Error(code.ServerErr).WithDetails(e))
				c.Abort()
			}
		}()

		c.Next()
	}
}
