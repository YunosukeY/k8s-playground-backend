package repository

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/util"
	"go.opentelemetry.io/otel/trace"
)

type dummyMailer struct {
	t trace.Tracer
}

func NewDummyMailer(t trace.Tracer) Mailer {
	return dummyMailer{t}
}

func (m dummyMailer) Send(ctx context.Context, mail model.Mail) error {
	_, span := m.t.Start(ctx, util.FuncName())
	defer span.End()

	return nil
}
