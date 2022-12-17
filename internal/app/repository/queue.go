package repository

import (
	"context"
	"encoding/json"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/trace"
)

const Topic = "mail"

func NewWriter() *kafka.Writer {
	addr := util.GetKafkaAddress()

	// check connection
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, Topic, 0)
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
	conn.Close()

	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    Topic,
		Balancer: &kafka.LeastBytes{},
	}
	return w
}

type Queue interface {
	Push(ctx context.Context, mail model.Mail) error
}

type queue struct {
	t trace.Tracer
	w *kafka.Writer
}

func NewQueue(t trace.Tracer, w *kafka.Writer) Queue {
	return queue{t, w}
}

func (q queue) Push(ctx context.Context, mail model.Mail) error {
	child, span := q.t.Start(ctx, util.FuncName())
	defer span.End()

	value, err := json.Marshal(mail)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}
	msg := kafka.Message{
		Key:   []byte("key"),
		Value: value,
	}
	err = q.w.WriteMessages(child, msg)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}
