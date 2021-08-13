package middleware

import (
	"self-discipline/global"
	"self-discipline/model/common/response"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Rate() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second*1), global.CONFIG.System.MaxBurstSize)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			response.FailWithMessage("Too Many Requests", c)
			c.Abort()
			return
		}
		// 继续往下处理
		c.Next()
	}
}
