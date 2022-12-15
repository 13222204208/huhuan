package initialize

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	fmt.Println("连接redis", redisCfg)
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	fmt.Println("客户端", client)
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接redis失败")
		global.GVA_LOG.Error("redis 连接失败, err:", zap.Error(err))
	} else {
		fmt.Println("连接redis成功")
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
