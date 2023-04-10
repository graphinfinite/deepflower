package telegram

import (
	"context"
	"deepflower/internal/model"
	"deepflower/internal/observer"
	"net/http"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"
)

type TelegramBot struct {
	Bot *tgbotapi.BotAPI
	log *zerolog.Logger
}

func NewBot(token string, buffer int, client http.Client, debug bool, logger *zerolog.Logger) (TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(token, &client)
	bot.Debug = debug
	bot.Buffer = buffer
	if err != nil {
		return TelegramBot{}, err
	}
	return TelegramBot{Bot: bot, log: logger}, nil
}

func (t *TelegramBot) StartReceiveUpdates(offset, limit, timeout int, outChan chan observer.Event) {
	t.log.Info().Msgf("Authorized on account %s", t.Bot.Self.UserName)
	updateschan, err := t.Bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset:  offset,
		Limit:   limit,
		Timeout: timeout})
	if err != nil {
		t.log.Panic().Msgf("StartReceiveUpdates/GetUpdatesChan %s", err.Error())
	}

	var event observer.Event
	for upd := range updateschan {
		usertg := model.UserTelegram{
			TgId:           upd.Message.From.ID,
			TgChatId:       upd.Message.Chat.ID,
			TgFirstName:    upd.Message.From.FirstName,
			TgLastName:     upd.Message.From.LastName,
			TgUserName:     upd.Message.From.UserName,
			TgLanguageCode: upd.Message.From.LanguageCode,
		}
		var message string
		var msg tgbotapi.MessageConfig

		switch upd.Message.Text {
		case "/start":
			message = "send '/auth' for registration"
			msg = tgbotapi.NewMessage(upd.Message.Chat.ID, message)
			t.Bot.Send(msg)

		case "/auth":
			event.Topic = "registration_from_tg"
			event.Payload = usertg
			outChan <- event
		}
		// end switch
	}
}

// var ErrAuthUserAlreadyExist *ErrAuthUserAlreadyExist
// usepas, err := t.Auth.RegistrationFromTg(context.Background(), usertg)
// switch {
// case errors.As(err, &ErrAuthUserAlreadyExist):
// 	message = fmt.Sprintf("Glad to see you here again, %s!", usepas.Username)
// case err != nil:
// 	t.log.Err(err).Msg("TelegramBotMessageReader ")
// 	message = "I'm broke. Sorry"
// default:
// 	message = fmt.Sprintf("Success registration!\n Username: %s \nPassword: %s", usepas.Username, usepas.Password)
// 	t.log.Info().Msgf("success registration. username: %s", usepas.Username)
// }
// msg = tgbotapi.NewMessage(upd.Message.Chat.ID, message)
// t.Bot.Send(msg)

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
