package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimit() gin.HandlerFunc {
	// 每100ms填充一个令牌，令牌桶容量为10000
	bucket := ratelimit.NewBucket(time.Microsecond*100, int64(10000))
	return func(c *gin.Context) {
		// 如果获取不到令牌就中断本次请求返回
		if bucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusOK, gin.H{
				"msg": "rate limit...",
			})
			c.Abort()
			return
		}
	}
}
