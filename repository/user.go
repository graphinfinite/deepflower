package repository

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type UserModel struct {
}

type UserStorage struct {
	Db *sql.DB
	L  *zerolog.Logger
}

func (u *UserStorage) GetUserByTgId(tgId string) {
}

func (u *UserStorage) UpdateUser(tgId string) {
}

func NewUserStorage(dbpool *sql.DB, logger *zerolog.Logger) UserStorage {
	return UserStorage{Db: dbpool, L: logger}
}
