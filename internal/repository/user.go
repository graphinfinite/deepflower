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
	err := s.Db.QueryRow(`SELECT * FROM users WHERE tgId = $1`, tgId).Scan(&user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, NewErrUserNotFound(fmt.Sprintf("user with telegramm id: %d not found", tgId), err)
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
	query := `INSERT INTO users(tgId, tgChatId, tgUserName, tgFirstName, tgLastName, tgLanguageCode, hashedPassword, username, active, createdAt) VALUES (?,?,?,?,?,?,?,?,?,?) returning id;`
	err := s.Db.QueryRow(query, u.TgId, u.TgChatId, u.TgUserName, u.TgFirstName, u.TgLastName, u.TgLanguageCode, u.HashedPassword, u.Username, false, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
