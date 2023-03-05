package controllers

import (
	h "deepflower/internal/helpers"
	"deepflower/internal/model"
	"deepflower/internal/usecase"
	"errors"
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type TelegramBot struct {
	Bot         *tgbotapi.BotAPI
	Logger      *zerolog.Logger
	Authusecase AuthUsecaseInterface
}

func NewBot(debug bool, client *http.Client, logger *zerolog.Logger, authusecase AuthUsecaseInterface) (TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(viper.GetString("telegram.token"), client)
	bot.Debug = debug
	if err != nil {
		return TelegramBot{}, err
	}
	return TelegramBot{Bot: bot, Logger: logger, Authusecase: authusecase}, nil
}

func (t *TelegramBot) TelegramBotMessageReader(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var upd tgbotapi.Update
	if err := h.DecodeJSONBody(w, r, &upd); err != nil {
		fmt.Println(err)
	}
	u := model.UserTelegram{
		TgId:           upd.Message.From.ID,
		TgChatId:       upd.Message.Chat.ID,
		TgFirstName:    upd.Message.From.FirstName,
		TgLastName:     upd.Message.From.LastName,
		TgUserName:     upd.Message.From.UserName,
		TgLanguageCode: upd.Message.From.LanguageCode,
	}
	t.Logger.Printf("User data: %#v", u)
	// registration
	if upd.Message.Text == "/start" {
		var message string
		var ErrAuthUserAlreadyExist *usecase.ErrAuthUserAlreadyExist
		usepas, err := t.Authusecase.RegistrationFromTg(u)
		switch {
		case errors.As(err, &ErrAuthUserAlreadyExist):
			message = fmt.Sprintf("Glad to see you here again, %s!", usepas.Username)
		case err != nil:
			message = "I'm broke. Sorry"
		default:
			message = fmt.Sprintf("Success registration!\n Username: %s \nPassword: %s", usepas.Username, usepas.Password)
		}
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
		t.Bot.Send(msg)
	}
	w.WriteHeader(http.StatusOK)
}
