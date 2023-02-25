package repository

import (
	"database/sql"
	"deepflower/model"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
)

type UserStorage struct {
	Db *sql.DB
	L  *zerolog.Logger
}

func NewUserStorage(dbpool *sql.DB, logger *zerolog.Logger) UserStorage {
	return UserStorage{Db: dbpool, L: logger}
}

// если пользователь не найден возвращает ошибку UserNotFoundStorageError
// db.error NewErrStoreUnknow
func (s *UserStorage) GetUserByTgId(tgId int) (model.User, error) {
	//TODO

	query := `SELECT username FROM user WHERE tgId = ?`
	user := model.User{}
	err := s.Db.QueryRow(query, tgId).Scan(&user.Username)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return user, NewErrUserNotFound(fmt.Sprintf("user with telegramm id: %d not found", user.TgId), err)

		default:
			return user, err
		}
	}
	return user, nil
}

func (s *UserStorage) CreateUser(tgId int, chatId int64, TgUserName, TgFirstName, TgLastName, TgLanguageCode, hash, newusername string) (int, error) {
	var id int
	query := `INSERT INTO user(tgId, chatId, TgUserName, tgFirstName, tgLastName, tgLanguageCode, hashedPassword, username) VALUES(?,?,?,?,?,?) returning id; `
	err := s.Db.QueryRow(query, tgId, chatId, TgUserName, TgFirstName, TgLastName, TgLanguageCode, hash, newusername).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
