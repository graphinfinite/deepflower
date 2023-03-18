package controllers

import (
	"deepflower/internal/model"
	"net/http"

	"github.com/rs/zerolog"
)

type UserController struct {
	Uc UserUCInterface
	L  *zerolog.Logger
}

func NewUserController(uc UserUCInterface, logger *zerolog.Logger) UserController {
	return UserController{Uc: uc, L: logger}
}

func (u *UserController) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(ContextUserIdKey).(string)
	user, err := u.Uc.GetUserInfo(userId)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "user info", user)
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := DecodeJSONBody(w, r, &user); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	userUpdated, err := u.Uc.UpdateUser(user)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "user successfully updated", userUpdated)
}
