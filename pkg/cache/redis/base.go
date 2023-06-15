package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vominhtrungpro/config"
)

const (
	redisClusterMode    = "cluster"
	redisStandaloneMode = "standalone"
)

type (
	Client interface {
		Get(ctx context.Context, key string) ([]byte, error)
		Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
		Del(ctx context.Context, keys ...string) error
		Close() error
		Ping(ctx context.Context) error
	}
)

func NewClient2(cfg *config.RedisConfig) (Client, error) {
	if cfg.Mode == redisStandaloneMode {
		return NewRedisClient(&cfg.Standalone)
	}
	return NewRedisClient(&cfg.Standalone)
	//return NewRedisClusterClient(&cfg.Cluster)
}

func NewClient(cfg *config.RedisConfig) (redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return *rdb, nil
}
