package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Item struct {
}

func (item *Item) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("item中间件开始")

		// before request

		c.Next()

		// after request
		fmt.Println("item中间件开始")
	}
}
