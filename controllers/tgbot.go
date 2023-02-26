package controllers

import (
	h "deepflower/helpers"
	"deepflower/usecase"
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
	Authusecase *usecase.AuthUsecase
}

func NewBot(debug bool, client *http.Client, logger *zerolog.Logger, authusecase *usecase.AuthUsecase) (TelegramBot, error) {
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

	tgId := upd.Message.From.ID
	chatId := upd.Message.Chat.ID
	firstName := upd.Message.From.FirstName
	lastName := upd.Message.From.LastName
	userName := upd.Message.From.UserName
	languageCode := upd.Message.From.LanguageCode

	t.Logger.Printf("user telegrammId %s ChatId %s firstname %s lastname %s username %s languageCode %s",
		tgId, chatId, firstName, lastName, userName, languageCode)

	// registration
	if upd.Message.Text == "/start" {
		var message string
		var ErrAuthUserAlreadyExist usecase.ErrAuthUserAlreadyExist
		usepas, err := t.Authusecase.RegistrationFromTg(tgId, chatId, userName, firstName, lastName, languageCode)
		if err != nil {
			if errors.Is(err, ErrAuthUserAlreadyExist) {
				message = fmt.Sprintf("Glad to see you here again, %s!", usepas.Username)
			} else {
				message = "I'm broke. Sorry"
				t.Logger.Printf("some problem with registration: %s", err.Error())

			}
		} else {
			message = fmt.Sprintf("Success Registration. Username: %s Password: %s", usepas.Username, usepas.Password)
		}
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
		t.Bot.Send(msg)
	}
	w.WriteHeader(http.StatusOK)
}
