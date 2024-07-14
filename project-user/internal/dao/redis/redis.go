/*
@author: NanYan
*/
package redis

import (
	"carrygpc.com/project-user/config"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"log"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(config.C.InitRedisConf())
	Rc = &RedisCache{rdb: rdb}
	result, err := Rc.rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}

func (r *RedisCache) Put(key string, value any, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := r.rdb.Set(ctx, key, value, expire).Err()
	if err != nil {
		log.Println("redis set error:", err)
		return err
	}
	return nil
}

func (r *RedisCache) Get(key string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *RedisCache) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return r.rdb.Del(ctx, key).Err()
}
