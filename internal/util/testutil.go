package util

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewTestDB(name string) (*gorm.DB, func()) {
	if err := loadEnv(); err != nil {
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
