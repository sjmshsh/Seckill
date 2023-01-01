package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	redis2 "github.com/sjmshsh/IM/db/redis"
	"github.com/sjmshsh/IM/rabbitmq"
	"github.com/sjmshsh/IM/request"
	"github.com/sjmshsh/IM/response"
	"log"
	"net/http"
)

func Seckill(request *request.SeckillRequest) *response.Response {
	// !!!!!!前提是我们已经做好了缓存预热，所以我们这里就直接把缓存写到redis里面去了
	// 1. 获取用户ID
	userId := request.UserId
	// 2. 获取产品ID
	productId := request.ProductId
	// 3. 执行Lua脚本
	var seckill = redis.NewScript(`
local productId = ARGV[1]
local userId = ARGV[2]
local stockKey = KEYS[1]..productId
local orderKey = KEYS[2]..productId
local res = redis.call('get', stockKey)
if (tonumber(res, 10) <= 0) then
    return 1
end
redis.call('incrby', stockKey, -1)
redis.call('sadd', orderKey, userId)
return 0
`)
	ctx := context.Background()
	keys := []string{"seckill:stock:", "seckill:order:"}
	values := []interface{}{productId, userId}
	resInterface, err := seckill.Run(ctx, redis2.Rdb, keys, values...).Result()
	res := resInterface.(int64)
	if err != nil {
		log.Println(err)
		log.Println("出现了严重的错误！！！")
	}
	if res == 1 {
		return &response.Response{
			Status: http.StatusOK,
			Msg:    "用户秒杀失败，库存不足",
		}
	}
	if res == 2 {
		return &response.Response{
			Status: http.StatusOK,
			Msg:    "用户秒杀失败，您已经秒杀过了，不可用继续秒杀",
		}
	}
	// 到这里就说明秒杀成功了
	// 这个时候我们可以进行异步下单
	mq := rabbitmq.NewRabbitMQTopics("seckill", "seckill-order")
	mq.PublishTopics(rabbitmq.Message{
		UserId:    userId,
		ProductId: productId,
	})
	return &response.Response{
		Status: http.StatusOK,
		Msg:    "用户秒杀成功",
	}
}

func InitSeckillConsumer() {
	mq := rabbitmq.NewRabbitMQTopics("seckill", "seckill-order")
	go mq.ConsumeTopics()
}
