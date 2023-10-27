package rs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/vildan-valeev/melvad_test/internal/domain"
	"github.com/vildan-valeev/melvad_test/internal/infra"
	"strconv"
)

func UpdateUserInCache(ctx context.Context, rdb infra.RedisCache, u domain.User) error {
	id := strconv.FormatInt(u.ID, 10) // тупое решение...

	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, id, "id", u.ID)
		rdb.HSet(ctx, id, "name", u.Name)
		rdb.HSet(ctx, id, "age", u.Age)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
