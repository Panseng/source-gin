package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"practise-code/config"
)

type Redis struct {
	Client *redis.Client
}

func InitRedis(cfg config.Redis, logger *zap.SugaredLogger) Redis{
	r := Redis{}
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
		Password: cfg.Password,
		DB: cfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil{
		logger.Errorw("redis connect ping error: " + err.Error())
		r.Client = nil
		return r
	}
	logger.Infof("redis connect success, ping is %+v", pong)

	r.Client = client
	return r
}
