package controllers

import (
	"net/http"

	"github.com/rs/zerolog"
)

type AuthController struct {
	Uc AuthUCInterface
	L  *zerolog.Logger
}

func NewAuthController(uc AuthUCInterface, logger *zerolog.Logger) AuthController {
	return AuthController{Uc: uc, L: logger}

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
	token, err := auth.Uc.Login(r.Context(), u.Username, u.Password)
	if err != nil {
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "token successfully generated", newSignInResponse(token))
}
