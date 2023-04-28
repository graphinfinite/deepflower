package repository

import (
	"context"
	"database/sql"
	"deepflower/internal/model"
	"deepflower/pkg/postgres"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type UserStorage struct {
	Db *postgres.PG
}

func NewUserStorage(db *postgres.PG) *UserStorage {
	return &UserStorage{Db: db}
}

// if user not found -> UserNotFoundStorageError
func (s *UserStorage) GetUserByTgId(ctx context.Context, tgId int) (model.User, error) {
	tx := s.Db.ExtractTx(ctx)
	q := `SELECT * FROM users WHERE tgId = $1`
	user := model.User{}
	err := tx.GetContext(ctx, &user, q, tgId)
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
	tx := s.Db.ExtractTx(ctx)
	q := `SELECT * FROM users WHERE username = $1`
	user := model.User{}
	err := tx.GetContext(ctx, &user, q, username)
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

// if user not found -> UserNotFoundStorageError
func (s *UserStorage) GetUserById(ctx context.Context, userId string) (model.User, error) {
	tx := s.Db.ExtractTx(ctx)
	q := `SELECT * FROM users WHERE id = $1`
	user := model.User{}
	err := tx.GetContext(ctx, &user, q, userId)
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
	tx := s.Db.ExtractTx(ctx)
	var id string
	query := `
	INSERT INTO users 
	(tgId, tgChatId, tgUserName, tgFirstName,
	tgLastName, tgLanguageCode, hashedPassword, username, active) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
	err := tx.GetContext(ctx, &id, query, u.TgId, u.TgChatId, u.TgUserName, u.TgFirstName, u.TgLastName, u.TgLanguageCode, u.HashedPassword, u.Username, false)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *UserStorage) SubtractEnergy(ctx context.Context, userId string, energy uint64) error {
	tx := s.Db.ExtractTx(ctx)
	query1 := `UPDATE users SET energy=energy-$1 WHERE id=$2;`

	_, err := tx.ExecContext(ctx, query1, energy, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) AddEnergy(ctx context.Context, userId string, energy uint64) error {
	tx := s.Db.ExtractTx(ctx)
	query1 := `UPDATE users SET energy=energy+$1 WHERE id=$2;`
	_, err := tx.ExecContext(ctx, query1, energy, userId)
	if err != nil {
		return err
	}
	return nil
}
