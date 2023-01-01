package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sjmshsh/IM/conf"
)

var RCtx = context.Background() // 全局Redis ctx
var Rdb *redis.Client           // 全局Redis DB

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,     // redis地址
		Password: conf.RedisPassword, // redis密码，没有则留空
		DB:       conf.RedisDbName,   // 默认数据库（不指定默认是0）
	})
}