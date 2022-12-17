package repository

import (
	"context"
	"fmt"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type dummyCache struct {
	t trace.Tracer
	m map[string]string
}

func NewDummyCache(t trace.Tracer) Cache {
	m := map[string]string{}
	return dummyCache{t, m}
}

func (c dummyCache) Get(ctx context.Context, key string) (string, error) {
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

func (c dummyCache) Set(ctx context.Context, key string, value string) error {
	_, span := c.t.Start(ctx, util.FuncName())
	defer span.End()

	c.m[key] = value

	return nil
}

func (c dummyCache) Delete(ctx context.Context, key string) error {
	_, span := c.t.Start(ctx, util.FuncName())
	defer span.End()

	delete(c.m, key)

	return nil
}
