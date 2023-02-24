package repository

import (
	"context"
	"database/sql"
	"deepflower/model"
)

func NewSqlitePool(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
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
