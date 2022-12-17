package repository

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/util"
	"go.opentelemetry.io/otel/trace"
)

type dummyQueue struct {
	t trace.Tracer
}

func NewDummyQueue(t trace.Tracer) Queue {
	return dummyQueue{t}
}

func (q dummyQueue) Push(ctx context.Context, mail model.Mail) error {
	_, span := q.t.Start(ctx, util.FuncName())
	defer span.End()

	return nil
}
