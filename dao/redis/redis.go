package redis

import (
	"gin_web_demo/setting"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(config *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         config.Host + ":" + strconv.Itoa(config.Port),
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return nil
	}
	return nil
}

// Close 关闭MySQL连接
func Close() {
	_ = client.Close()
}
