package migration

import (
	"path/filepath"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

type Op int

const (
	Up Op = iota
	Down
)

type Type int

const (
	Table Type = iota
	Record
)

func mig(op Op, ty Type) {
	path, err := util.SchemaPath()
	if err != nil {
		panic(err)
	}
	source := filepath.Join("file://", path)
	if ty == Table {
		source = filepath.Join(source, "schema")
	} else {
		source = filepath.Join(source, "data")
	}
	db := "mysql://" + util.GetConnectionString()
	m, err := migrate.New(source, db)
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}

	if op == Up {
		err = m.Up()
	} else {
		err = m.Down()
	}
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Info().Err(err).Msg("")
		} else {
			log.Panic().Err(err).Msg("")
			panic(err)
		}
	}
}

func UpSchema() {
	mig(Up, Table)
}

func DownSchema() {
	mig(Down, Table)
}

func UpTestData() {
	mig(Up, Record)
}

func DownTestData() {
	mig(Down, Record)
}
