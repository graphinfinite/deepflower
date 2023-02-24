package controllers

import (
	"deepflower/usecase"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type AuthController struct {
	Uc *usecase.AuthUsecase
	L  *zerolog.Logger
}

func NewAuthController(uc *usecase.AuthUsecase) AuthController {
	return AuthController{Uc: uc}

}

func (auth *AuthController) RedirectToTelegram(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, viper.GetString("telegram.boturl"), http.StatusPermanentRedirect)
}
