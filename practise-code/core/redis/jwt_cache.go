package redis

import (
	"context"
	"time"
)

func (r Redis)SetRedisJWT(jwt string, userName string, timer time.Duration) error {
	return r.Client.Set(context.Background(), userName, jwt, timer).Err()
}

func (r Redis)GetRedisJWT(userName string)(jwt string, err error)  {
	jwt, err = r.Client.Get(context.Background(), userName).Result()
	return jwt, err
}
