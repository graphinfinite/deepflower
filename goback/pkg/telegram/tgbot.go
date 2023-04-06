package telegram

import (
	"context"
	"deepflower/internal/model"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"
)

type Authorizater interface {
	RegistrationFromTg(ctx context.Context, tguser model.UserTelegram) (model.User, error)
}
type ErrAuthUserAlreadyExist struct {
	msg string
	err error
}

func NewErrAuthUserAlreadyExist(msg string, err error) *ErrAuthUserAlreadyExist {
	return &ErrAuthUserAlreadyExist{msg: msg, err: err}
}
func (err ErrAuthUserAlreadyExist) Error() string {
	return err.msg
}
func (err ErrAuthUserAlreadyExist) Unwrap() error {
	return err.err
}

type TelegramBot struct {
	Bot  *tgbotapi.BotAPI
	log  *zerolog.Logger
	Auth Authorizater
}

func NewBot(token string, debug bool, client http.Client, logger *zerolog.Logger, auth Authorizater) (TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(token, &client)
	bot.Debug = debug
	if err != nil {
		return TelegramBot{}, err
	}
	return TelegramBot{Bot: bot, log: logger, Auth: auth}, nil
}

func (t *TelegramBot) StartReceiveUpdates() {
	t.log.Info().Msgf("Authorized on account %s", t.Bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updateschan, err := t.Bot.GetUpdatesChan(tgbotapi.UpdateConfig{})
	if err != nil {
		t.log.Panic().Msgf("StartReceiveUpdates/GetUpdatesChan %s", err.Error())
	}
	for upd := range updateschan {
		usertg := model.UserTelegram{
			TgId:           upd.Message.From.ID,
			TgChatId:       upd.Message.Chat.ID,
			TgFirstName:    upd.Message.From.FirstName,
			TgLastName:     upd.Message.From.LastName,
			TgUserName:     upd.Message.From.UserName,
			TgLanguageCode: upd.Message.From.LanguageCode,
		}
		//t.log.Debug().Msgf("user data from tg: %#v", usertg)
		if upd.Message.Text == "/start" {
			var message string
			var ErrAuthUserAlreadyExist *ErrAuthUserAlreadyExist
			usepas, err := t.Auth.RegistrationFromTg(context.Background(), usertg)
			switch {
			case errors.As(err, &ErrAuthUserAlreadyExist):
				message = fmt.Sprintf("Glad to see you here again, %s!", usepas.Username)
			case err != nil:
				t.log.Err(err).Msg("TelegramBotMessageReader ")
				message = "I'm broke. Sorry"
			default:
				message = fmt.Sprintf("Success registration!\n Username: %s \nPassword: %s", usepas.Username, usepas.Password)
				t.log.Info().Msgf("success registration. username: %s", usepas.Username)
			}
			msg := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
			t.Bot.Send(msg)
		}
	}
}

func (t *TelegramBot) SendMessages(ctx context.Context, chatIds []string, msg string) error {
	for _, chatId := range chatIds {
		id, err := strconv.ParseInt(chatId, 0, 64)
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(id, msg)
		_, err = t.Bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil

}
