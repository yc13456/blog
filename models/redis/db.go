package redis

import (
	"errors"
	"github.com/go-redis/redis"
)

type redisDB struct {
	client *redis.Client
}

func NewRedisDB() *redisDB {
	return &redisDB{
		client: redis.NewClient(&redis.Options{
			Addr:     "192.168.1.106:6379", //连接地址
			Password: "",               //连接密码
			DB:       0,                //默认连接库
			PoolSize: 100,              //连接池大小
		}),
	}
}

func (c *redisDB) Connect() error{
	_, err := c.client.Ping().Result()
	if err != nil {
		return errors.New("redis connect error")
	}
	return nil
}