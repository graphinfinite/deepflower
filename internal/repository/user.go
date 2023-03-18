package repository

import (
	"database/sql"
	"deepflower/internal/model"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserStorage struct {
	Db *sqlx.DB
}

func NewUserStorage(dbpool *sqlx.DB) UserStorage {
	return UserStorage{Db: dbpool}
}

// return user
// if user not found -> UserNotFoundStorageError
func (s *UserStorage) GetUserByTgId(tgId int) (model.User, error) {
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

func (s *UserStorage) GetUserByUsername(username string) (model.User, error) {
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

func (s *UserStorage) UpdateUser(m model.User) (model.User, error) {
	//q := `UPDATE users SET username = $1`

	return model.User{}, nil
}

func (s *UserStorage) GetUserById(id string) (model.User, error) {
	q := `SELECT * FROM users WHERE id = $1`
	user := model.User{}
	err := s.Db.Get(&user, q, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, NewErrUserNotFound(fmt.Sprintf("user with username: %s not found", id), err)
		} else {
			return user, err
		}
	}
	return user, nil
}

func (s *UserStorage) CreateUser(u model.User) (userId int, e error) {
	var id int
	fmt.Printf("USER  %#v \n", u)
	query := `INSERT INTO users (tgId, tgChatId, tgUserName, tgFirstName, tgLastName, tgLanguageCode, hashedPassword, username, active, createdAt) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id`
	err := s.Db.QueryRow(query, u.TgId, u.TgChatId, u.TgUserName, u.TgFirstName, u.TgLastName, u.TgLanguageCode, u.HashedPassword, u.Username, false, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
