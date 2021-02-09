// 500错误统一监听处理
package middleware

import (
	"fmt"
	"gin/output"
	"github.com/gin-gonic/gin"
)

type Panic struct {
}

// 统一错误处理
// TODO @2021-02-09 后续可使用 CustomRecovery 方法，待gin框架发布新版
func (panic *Panic) Handler() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("自定义Panicing %s\r\n", e)
				output.Error(c, e)
				c.Abort()
			}
		}()

		c.Next()
	}
}