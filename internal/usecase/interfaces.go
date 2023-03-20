package usecase

import (
	"deepflower/internal/model"
)

type UserStorageInterface interface {
	CreateUser(u model.User) (string, error)
	GetUserByTgId(tgId int) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
	GetUserById(userId string) (model.User, error)
	UpdateUser(model.User) (model.User, error)
}

type DreamStorageInterface interface {
	CreateDream(name, info, location, creater string) (model.Dream, error)
	GetAllUserDreams(userId string) ([]model.Dream, error)
	GetDreamById(dreamId string) (model.Dream, error)
	DeleteUserDream(dreamId string) error
	UpdateUserDream(dreamId string, patchDream map[string]interface{}) (model.Dream, error)
}
