package auth

import (
	"context"
	"fmt"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
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
	get(ctx context.Context, key string) (string, error)
	set(ctx context.Context, key string, value string) error
	delete(ctx context.Context, key string) error
}

type cache struct {
	t trace.Tracer
	r *redis.Client
}

func NewCache(t trace.Tracer, r *redis.Client) Cache {
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

type dummyCache struct {
	t trace.Tracer
	m map[string]string
}

func NewDummyCache(t trace.Tracer) Cache {
	m := map[string]string{}
	return dummyCache{t, m}
}

func (c dummyCache) get(ctx context.Context, key string) (string, error) {
	_, span := c.t.Start(ctx, util.FuncName())
	defer span.End()

	v, ok := c.m[key]
	if !ok {
		err := fmt.Errorf("no such key: %s", key)
		log.Error().Err(err).Msg("")
		return "", err
	}
	return v, nil
}

func (c dummyCache) set(ctx context.Context, key string, value string) error {
	_, span := c.t.Start(ctx, util.FuncName())
	defer span.End()

	c.m[key] = value

	return nil
}

func (c dummyCache) delete(ctx context.Context, key string) error {
	_, span := c.t.Start(ctx, util.FuncName())
	defer span.End()

	delete(c.m, key)

	return nil
}
