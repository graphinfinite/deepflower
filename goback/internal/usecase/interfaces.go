package usecase

import (
	"context"
	"deepflower/internal/model"
)

type UserStorageInterface interface {
	CreateUser(ctx context.Context, u model.User) (string, error)
	GetUserByTgId(ctx context.Context, tgId int) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	GetUserById(ctx context.Context, userId string) (model.User, error)
	UpdateUser(context.Context, model.User) (model.User, error)
}

type DreamStorageInterface interface {
	CreateDream(ctx context.Context, name, info, location, creater string) (model.Dream, error)
	GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error)
	GetDreamById(ctx context.Context, dreamId string) (model.Dream, error)
	DeleteUserDream(ctx context.Context, dreamId string) error
	UpdateUserDream(ctx context.Context, dreamId string, patchDream map[string]interface{}) (model.Dream, error)
}
