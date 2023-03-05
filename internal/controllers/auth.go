package controllers

import (
	h "deepflower/internal/helpers"
	"deepflower/internal/model"
	"encoding/json"
	"net/http"

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
}

func newResponse(status, msg string) *response {
	return &response{
		Status: status,
		Msg:    msg,
	}
}

func NewAuthController(uc AuthUsecaseInterface, logger *zerolog.Logger) AuthController {
	return AuthController{Uc: uc, L: logger}

}

func (auth *AuthController) RedirectToTelegram(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, viper.GetString("telegram.boturl"), http.StatusPermanentRedirect)
}

// Login ----->
type loginUser struct {
	Username string
	Password string
}

type signInResponse struct {
	*response
	Token string `json:"token,omitempty"`
}

func newSignInResponse(status, msg, token string) *signInResponse {
	return &signInResponse{
		&response{
			Status: status,
			Msg:    msg,
		},
		token,
	}
}

// TODO
func (auth *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	u := loginUser{}
	j := json.NewEncoder(w)
	if err := h.DecodeJSONBody(w, r, &u); err != nil {
		j.Encode(newResponse(STATUS_ERROR, ""))
	}

	token, err := auth.Uc.Login(u.Username, u.Password)
	if err != nil {
		j.Encode(newResponse(STATUS_ERROR, ""))
	}

	signInRespons := newSignInResponse(STATUS_OK, "", token)
	if err = j.Encode(signInRespons); err != nil {
		j.Encode(newResponse(STATUS_ERROR, ""))
	}
	w.WriteHeader(http.StatusOK)
}
