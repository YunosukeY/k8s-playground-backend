package repository

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

func NewRedis() *redis.Client {
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
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
	Delete(ctx context.Context, key string) error
}

type cache struct {
	t trace.Tracer
	r *redis.Client
}

func NewCache(t trace.Tracer, r *redis.Client) Cache {
	return cache{t, r}
}

func (c cache) Get(ctx context.Context, key string) (string, error) {
	v, err := c.r.Get(ctx, key).Result()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return v, err
}

func (c cache) Set(ctx context.Context, key string, value string) error {
	err := c.r.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}

func (c cache) Delete(ctx context.Context, key string) error {
	err := c.r.Del(ctx, key).Err()
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}
