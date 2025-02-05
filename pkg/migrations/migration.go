package migrations

import (
	"backend/pkg/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(config config.Config) {
	m, err := migrate.New(
		"file://pkg/migrations/query",
		config.DBConnectionString,
	)

	if err != nil {
		log.Println("Error while creating migration instance")
		panic(err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Println("Error while migrating up")
		panic(err)
	}

	log.Println("Migration success")
}
