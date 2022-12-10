package mail

import (
	"context"
	"fmt"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type controller struct {
	t trace.Tracer
	q Queue
	m Mailer
}

func newController(t trace.Tracer, q Queue, m Mailer) controller {
	return controller{t, q, m}
}

func (c controller) handle() {
	fmt.Println("mailer started")
	for {
		child, span := c.t.Start(context.Background(), util.FuncName())

		m, err := c.q.pop(child)
		if err != nil {
			panic(err)
		}
		log.Debug().Interface("mail", m).Msg("")
		if err := c.m.send(child, *m); err != nil {
			panic(err)
		}

		span.End()
	}
}

func Run(dummy bool) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.With().Caller().Logger()
	gin.SetMode(gin.ReleaseMode)

	var c controller
	var shutdownProvider func()
	if !dummy {
		c, shutdownProvider = initializeController("mail")
	} else {
		c, shutdownProvider = initializeDummyController("mail")
	}
	defer shutdownProvider()

	go util.RunPodCommonHandler()
	c.handle()
}
