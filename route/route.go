package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sjmshsh/IM/api"
	"github.com/sjmshsh/IM/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RateLimit())
	r.GET("/seckill", api.Seckill)
	return r
}
