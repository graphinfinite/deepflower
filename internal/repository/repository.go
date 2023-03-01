package repository

import (
	"context"
	"database/sql"
	"deepflower/internal/model"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func MigrateDb(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}

func NewPostgresPool(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

type UserStorageInteface interface {
	CreateUser() error
	GetUserByTgId(ctx context.Context, tgId string) (model.User, error)
	UpdateUser() (model.User, error)
	DeleteUser() error
}
