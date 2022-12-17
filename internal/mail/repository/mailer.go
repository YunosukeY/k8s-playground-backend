package repository

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type Mailer interface {
	Send(ctx context.Context, mail model.Mail) error
}

type mailer struct {
	t    trace.Tracer
	auth smtp.Auth
	addr string
}

func NewMailer(t trace.Tracer) Mailer {
	username := util.GetParamString("MAIL_USER", "user")
	password := util.GetParamString("MAIL_PASS", "pass")
	auth := smtp.CRAMMD5Auth(username, password)
	addr := util.GetMailServerAddress()
	return mailer{t, auth, addr}
}

func (m mailer) Send(ctx context.Context, mail model.Mail) error {
	_, span := m.t.Start(ctx, util.FuncName())
	defer span.End()

	from := "test@example.com"

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", mail.To, *mail.Sub, *mail.Msg))
	err := smtp.SendMail(m.addr, m.auth, from, []string{mail.To}, msg)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}
