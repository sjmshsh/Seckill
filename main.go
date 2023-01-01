package main

import (
	"github.com/sjmshsh/IM/conf"
	"github.com/sjmshsh/IM/db/mysql"
	"github.com/sjmshsh/IM/db/redis"
	"github.com/sjmshsh/IM/route"
	"github.com/sjmshsh/IM/service"
)

func main() {
	r := route.NewRouter()
	r.Run(conf.HttpPort)
}

func init() {
	conf.Init("./conf/config.ini")
	redis.Init()
	mysql.Init()
	service.InitSeckillConsumer()
	// mongo.Init()
}
