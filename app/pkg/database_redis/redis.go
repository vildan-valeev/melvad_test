package database_redis

import (
	"context"
	"net"

	"github.com/redis/go-redis/v9"
)

type cache struct {
	client *redis.Client

	redisHost string
	redisPort string
}

func New(host, port string) *cache {
	return &cache{
		redisHost: host,
		redisPort: port,
	}
}

func (rs *cache) Open(ctx context.Context) error {
	rs.client = redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379",
		Addr:     net.JoinHostPort(rs.redisHost, rs.redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}

func (rs *cache) Close() error {
	return nil
}

func (rs *cache) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return rs.client.HSet(ctx, key, values)
}

func (rs *cache) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return rs.client.Pipeline().Pipelined(ctx, fn)
}

func (rs *cache) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	return rs.client.HGetAll(ctx, key)
}

//func (rs cache) Scan(dest interface{}) error {
//	return rs.client.Scan()
//}
