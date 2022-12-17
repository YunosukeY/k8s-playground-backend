package mail

import (
	"context"
	"encoding/json"
	"time"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/app/repository"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/trace"
)

func NewReader() *kafka.Reader {
	addr := util.GetKafkaAddress()

	// check connection
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, repository.Topic, 0)
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
	conn.Close()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{addr},
		Topic:     repository.Topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	return r
}

type Queue interface {
	pop(ctx context.Context) (*model.Mail, error)
}

type queue struct {
	t trace.Tracer
	r *kafka.Reader
}

func NewQueue(t trace.Tracer, r *kafka.Reader) Queue {
	return queue{t, r}
}

func (q queue) pop(ctx context.Context) (*model.Mail, error) {
	child, span := q.t.Start(ctx, util.FuncName())
	defer span.End()

	m, err := q.r.ReadMessage(child)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	var mail model.Mail
	if err := json.Unmarshal(m.Value, &mail); err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}
	return &mail, nil
}

type dummyQueue struct {
	t trace.Tracer
}

func NewDummyQueue(t trace.Tracer) Queue {
	return dummyQueue{t}
}

func (q dummyQueue) pop(ctx context.Context) (*model.Mail, error) {
	_, span := q.t.Start(ctx, util.FuncName())
	defer span.End()

	time.Sleep(time.Second * 10)

	sub := "title"
	msg := "msg"
	mail := model.Mail{To: "test@example.com", Sub: &sub, Msg: &msg}

	return &mail, nil
}
