package repository

import (
	"context"
	"database/sql"
	"deepflower/internal/model"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserStorage struct {
	Db *sqlx.DB
}

func NewUserStorage(dbpool *sqlx.DB) *UserStorage {
	return &UserStorage{Db: dbpool}
}

// if user not found -> UserNotFoundStorageError
func (s *UserStorage) GetUserByTgId(ctx context.Context, tgId int) (model.User, error) {
	q := `SELECT * FROM users WHERE tgId = $1`
	user := model.User{}
	err := s.Db.Get(&user, q, tgId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, NewErrUserNotFound(fmt.Sprintf("user with telegramm id: %d not found", tgId), err)
		} else {
			return user, err
		}
	}
	return user, nil
}

// if user not found -> UserNotFoundStorageError
func (s *UserStorage) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	q := `SELECT * FROM users WHERE username = $1`
	user := model.User{}
	err := s.Db.Get(&user, q, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, NewErrUserNotFound(fmt.Sprintf("user with username: %s not found", username), err)
		} else {
			fmt.Print(err)
			return user, err

		}
	}
	return user, nil
}

// не реализован
func (s *UserStorage) UpdateUser(ctx context.Context, m model.User) (model.User, error) {
	//q := `UPDATE users SET username = $1`

	return model.User{}, nil
}

// if user not found -> UserNotFoundStorageError
func (s *UserStorage) GetUserById(ctx context.Context, userId string) (model.User, error) {
	q := `SELECT * FROM users WHERE id = $1`
	user := model.User{}
	err := s.Db.Get(&user, q, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, NewErrUserNotFound(fmt.Sprintf("user with username: %s not found", userId), err)
		} else {
			return user, err
		}
	}
	return user, nil
}

func (s *UserStorage) CreateUser(ctx context.Context, u model.User) (userId string, e error) {
	var id string
	fmt.Printf("USER  %#v \n", u)
	query := `INSERT INTO users (tgId, tgChatId, tgUserName, tgFirstName, tgLastName, tgLanguageCode, hashedPassword, username, active) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
	err := s.Db.QueryRow(query, u.TgId, u.TgChatId, u.TgUserName, u.TgFirstName, u.TgLastName, u.TgLanguageCode, u.HashedPassword, u.Username, false).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
