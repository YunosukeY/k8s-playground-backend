package auth

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func newRedis() *redis.Client {
	op := util.GetRedisOptions()
	r := redis.NewClient(op)
	if err := redisotel.InstrumentTracing(r); err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}

	if err := r.Ping(context.Background()).Err(); err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}

	return r
}

type Cache interface {
	get(ctx context.Context, key string) (string, error)
	set(ctx context.Context, key string, value string) error
	delete(ctx context.Context, key string) error
}

type cache struct {
	t trace.Tracer
	r *redis.Client
}

func newCache(t trace.Tracer, r *redis.Client) cache {
	return cache{t, r}
}

func (c cache) get(ctx context.Context, key string) (string, error) {
	v, err := c.r.Get(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return v, err
}

func (c cache) set(ctx context.Context, key string, value string) error {
	err := c.r.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}

func (c cache) delete(ctx context.Context, key string) error {
	err := c.r.Del(ctx, key).Err()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}
