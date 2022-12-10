package mail

import (
	"context"
	"encoding/json"

	"github.com/YunosukeY/kind-backend/internal/app"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/trace"
)

func newReader() *kafka.Reader {
	addr := util.GetKafkaAddress()

	// check connection
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, app.Topic, 0)
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
	conn.Close()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{addr},
		Topic:     app.Topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	return r
}

type Queue interface {
	pop(ctx context.Context) (*app.Mail, error)
}

type queue struct {
	t trace.Tracer
	r *kafka.Reader
}

func newQueue(t trace.Tracer, r *kafka.Reader) Queue {
	return queue{t, r}
}

func (q queue) pop(ctx context.Context) (*app.Mail, error) {
	child, span := q.t.Start(ctx, util.FuncName())
	defer span.End()

	m, err := q.r.ReadMessage(child)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	var mail app.Mail
	if err := json.Unmarshal(m.Value, &mail); err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	return &mail, nil
}

type dummyQueue struct {
	t trace.Tracer
}

func newDummyQueue(t trace.Tracer) Queue {
	return dummyQueue{t}
}

func (q dummyQueue) pop(ctx context.Context) (*app.Mail, error) {
	_, span := q.t.Start(ctx, util.FuncName())
	defer span.End()

	sub := "title"
	msg := "msg"
	mail := app.Mail{To: "test@example.com", Sub: &sub, Msg: &msg}

	return &mail, nil
}
