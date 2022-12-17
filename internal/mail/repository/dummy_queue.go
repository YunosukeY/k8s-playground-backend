package repository

import (
	"context"
	"time"

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

func (q dummyQueue) Pop(ctx context.Context) (*model.Mail, error) {
	_, span := q.t.Start(ctx, util.FuncName())
	defer span.End()

	time.Sleep(time.Second * 10)

	sub := "title"
	msg := "msg"
	mail := model.Mail{To: "test@example.com", Sub: &sub, Msg: &msg}

	return &mail, nil
}
