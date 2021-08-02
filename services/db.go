package services

import (
	"database/sql"
	"embed"
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func SetupDB(migrations embed.FS) (*sql.DB, error) {
	// setup postgres
	pgConn, ok := os.LookupEnv("PG_CONN")
	if !ok {
		return nil, errors.New("PG_CONN not available")
	}
	db, err := sql.Open("postgres", pgConn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	d, err := iofs.New(migrations, "sql")
	if err != nil {
		return nil, err
	}
	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithInstance("iofs", d, "postgres", instance)
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}
	return db, nil
}
