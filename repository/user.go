package repository

import (
	"database/sql"
	"deepflower/model"
	"errors"
	"fmt"
)

type UserStorage struct {
	Db *sql.DB
}

func NewUserStorage(dbpool *sql.DB) UserStorage {
	return UserStorage{Db: dbpool}
}

// return user
// if user not found -> UserNotFoundStorageError
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

// return user id(int)
func (s *UserStorage) CreateUser(tgId int, tgchatId int64, tgUserName, tgFirstName, tgLastName, tgLanguageCode, hash, newusername string) (int, error) {
	var id int
	query := `INSERT INTO user(tgId, chatId, TgUserName, tgFirstName, tgLastName, tgLanguageCode, hashedPassword, username) VALUES(?,?,?,?,?,?) returning id; `
	err := s.Db.QueryRow(query, tgId, tgchatId, tgUserName, tgFirstName, tgLastName, tgLanguageCode, hash, newusername).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
