package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sjmshsh/IM/request"
	"github.com/sjmshsh/IM/service"
	"log"
	"strconv"
)

// Seckill /seckill?id=100
func Seckill(ctx *gin.Context) {
	idString := ctx.Query("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
	}
	req := &request.SeckillRequest{
		ProductId: id,
		UserId:    1,
	}
	log.Println(req)
	response := service.Seckill(req)
	ctx.JSON(response.Status, response.Msg)
}
