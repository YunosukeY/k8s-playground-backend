package util

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/rs/zerolog/log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewTestDB(name string) (*gorm.DB, func()) {
	if err := LoadEnv(); err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
	dsn := GetConnectionString()

	txdb.Register(name, "mysql", dsn)

	db, err := sql.Open(name, dsn)
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}

	return gormDB, func() { db.Close() }
}

func NewTestTracer() trace.Tracer {
	return sdktrace.NewTracerProvider().Tracer("")
}
