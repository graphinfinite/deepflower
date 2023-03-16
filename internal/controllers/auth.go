package controllers

import (
	"deepflower/internal/model"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type AuthController struct {
	Uc AuthUsecaseInterface
	L  *zerolog.Logger
}

type AuthUsecaseInterface interface {
	RegistrationFromTg(tguser model.UserTelegram) (model.User, error)
	Login(username, password string) (token string, err error)
	ValidateJwtToken(tokenString string) (bool, jwt.MapClaims, error)
}

func NewAuthController(uc AuthUsecaseInterface, logger *zerolog.Logger) AuthController {
	return AuthController{Uc: uc, L: logger}

}

func (auth *AuthController) RedirectToTelegram(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, viper.GetString("telegram.boturl"), http.StatusPermanentRedirect)
}

// Login ----->
type loginUser struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
type signInResponse struct {
	Token string `json:"token,omitempty"`
}

func newSignInResponse(token string) *signInResponse {
	return &signInResponse{
		Token: token,
	}
}

// TODO logger + ERROR handle
func (auth *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	u := loginUser{}
	if err := DecodeJSONBody(w, r, &u); err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	token, err := auth.Uc.Login(u.Username, u.Password)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "token successfully generated", newSignInResponse(token))
}
