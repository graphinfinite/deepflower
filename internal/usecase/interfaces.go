package usecase

import m "deepflower/internal/model"

type UserStorageInterface interface {
	CreateUser(u m.User) (int, error)
	GetUserByTgId(tgId int) (m.User, error)
	GetUserByUsername(username string) (m.User, error)
	GetUserById(id string) (m.User, error)
}
