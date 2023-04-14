package controllers

import (
	"context"
	"deepflower/internal/model"
	"deepflower/internal/observer"
	"deepflower/internal/usecase"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

type AuthController struct {
	Uc  AuthUCInterface
	log *zerolog.Logger
	bot BotInterface
}

func NewAuthController(uc AuthUCInterface, bot BotInterface, logger *zerolog.Logger) AuthController {
	return AuthController{Uc: uc, bot: bot, log: logger}

}

// Login ----->
type loginUser struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
type SignInResponse struct {
	Token string `json:"token,omitempty"`
}

func (auth *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	u := loginUser{}
	if err := DecodeJSONBody(w, r, &u); err != nil {
		auth.log.Err(err).Msg("Login ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	token, err := auth.Uc.Login(r.Context(), u.Username, u.Password)
	if err != nil {
		auth.log.Err(err).Msg("Login ")
		JSON(w, STATUS_ERROR, err.Error())
		return
	}
	JSONstruct(w, STATUS_OK, "token successfully generated", SignInResponse{Token: token})
}

func (auth *AuthController) Registration(event observer.Event) {
	ctx := context.Background()
	switch event.Topic {
	case "bot/registration":
		e, ok := event.Payload.(model.UserTelegram)
		if !ok {
			auth.log.Error().Msg("Registration/event/error payload type")
		}
		user, err := auth.Uc.RegistrationFromTg(ctx, e)
		if err != nil {
			auth.log.Err(err).Msg("Registration/RegistrationFromTg ")
		}
		var ErrAuthUserAlreadyExist *usecase.ErrAuthUserAlreadyExist
		var m string
		switch {
		case errors.As(err, &ErrAuthUserAlreadyExist):
			m = fmt.Sprintf("Glad to see you here again, %s!", user.Username)
		case err != nil:
			auth.log.Err(err).Msg("Registration ")
			m = "I'm broke. Sorry"
		default:
			m = fmt.Sprintf("Success registration!\n Username: %s \nPassword: %s", user.Username, user.Password)
			auth.log.Info().Msgf("success registration. username: %s", user.Username)
		}
		if err := auth.bot.SendMessage(ctx, e.TgChatId, m); err != nil {
			auth.log.Err(err).Msg("Registration/send message error ")

		}
	default:
		auth.log.Error().Msg("Registration/event/ unknow topic")
	}

}
