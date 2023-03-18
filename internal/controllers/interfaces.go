package controllers

import (
	"deepflower/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthUCInterface interface {
		RegistrationFromTg(tguser model.UserTelegram) (model.User, error)
		Login(username, password string) (token string, err error)
		ValidateJwtToken(tokenString string) (bool, jwt.MapClaims, error)
	}
	UserUCInterface interface {
		GetUserInfo(userId string) (user model.User, err error)
		UpdateUser(model.User) (user model.User, err error)
	}
	DreamUCInterface interface {
		CreateDream(name, info, location, creater string) (model.Dream, error)
		GetAllUserDreams(userId string) ([]model.Dream, error)
	}

	TaskUsecaseInterface interface {
	}
)
