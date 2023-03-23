package controllers

import (
	"context"
	"deepflower/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthUCInterface interface {
		RegistrationFromTg(ctx context.Context, tguser model.UserTelegram) (model.User, error)
		Login(ctx context.Context, username, password string) (token string, err error)
		ValidateJwtToken(ctx context.Context, tokenString string) (bool, jwt.MapClaims, error)
	}
	UserUCInterface interface {
		GetUserInfo(ctx context.Context, userId string) (user model.User, err error)
		UpdateUser(context.Context, model.User) (user model.User, err error)
	}
	DreamUCInterface interface {
		CreateDream(ctx context.Context, name, info, location string, creater string) (model.Dream, error)
		GetAllUserDreams(ctx context.Context, userId string) ([]model.Dream, error)
		UpdateUserDream(ctx context.Context, userId, dreamId string, dream map[string]interface{}) (model.Dream, error)
		DeleteUserDream(ctx context.Context, userId string, dreamId string) error
		AddEnergyToDream(ctx context.Context, userId, dreamId string, energy uint64) error
		PublishDream(ctx context.Context, userId, dreamId string) error
		SearchDreams(ctx context.Context, userId string,
			limit uint64, offset uint64, onlyMyDreams bool,
			order string, searchTerm string,
			sort string) ([]model.Dream, error)
	}

	TaskUsecaseInterface interface {
	}
)
