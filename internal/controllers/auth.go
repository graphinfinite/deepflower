package controllers

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type AuthController struct {
	Uc AuthUsecaseInterface
	L  *zerolog.Logger
}

func NewAuthController(uc AuthUsecaseInterface, logger *zerolog.Logger) AuthController {
	return AuthController{Uc: uc, L: logger}

}

func (auth *AuthController) RedirectToTelegram(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, viper.GetString("telegram.boturl"), http.StatusPermanentRedirect)
}
