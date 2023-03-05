package repository

import (
	"database/sql"
	"deepflower/internal/model"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
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
	user := model.User{}
	err := s.Db.QueryRow(`SELECT id,createdAt,updatedAt,username,hashedPassword,active,tgId,tgChatId,tgUserName,tgFirstName,tgLastName,tgLanguageCode FROM users WHERE tgId = $1`, tgId).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Username,
		&user.HashedPassword,
		&user.Active,
		&user.TgId,
		&user.TgChatId,
		&user.TgUserName,
		&user.TgFirstName,
		&user.TgLastName,
		&user.TgLanguageCode,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, NewErrUserNotFound(fmt.Sprintf("user with telegramm id: %d not found", tgId), err)
		} else {
			return user, err
		}
	}
	return user, nil
}

func (s *UserStorage) GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	err := s.Db.QueryRow(`SELECT id,createdAt,updatedAt,username, hashedPassword,active,tgId,tgChatId,tgUserName,tgFirstName,tgLastName,tgLanguageCode FROM users WHERE username = $1`, username).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Username,
		&user.HashedPassword,
		&user.Active,
		&user.TgId,
		&user.TgChatId,
		&user.TgUserName,
		&user.TgFirstName,
		&user.TgLastName,
		&user.TgLanguageCode,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, NewErrUserNotFound(fmt.Sprintf("user with username: %s not found", username), err)
		} else {
			return user, err
		}
	}
	return user, nil
}

// return user id(int)
func (s *UserStorage) CreateUser(u model.User) (int, error) {
	var id int
	fmt.Printf("USER  %#v \n", u)
	query := `INSERT INTO users (tgId, tgChatId, tgUserName, tgFirstName, tgLastName, tgLanguageCode, hashedPassword, username, active, createdAt) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id`
	err := s.Db.QueryRow(query, u.TgId, u.TgChatId, u.TgUserName, u.TgFirstName, u.TgLastName, u.TgLanguageCode, u.HashedPassword, u.Username, false, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
