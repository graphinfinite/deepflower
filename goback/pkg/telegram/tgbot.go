package telegram

import (
	"context"
	"deepflower/internal/model"
	"deepflower/internal/observer"
	"net/http"

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
			event.Topic = "bot/registration"
			event.Payload = usertg
			outChan <- event
		}
		// end switch
	}
}

func (t *TelegramBot) SendMessage(ctx context.Context, chatId int64, message string) error {

	msg := tgbotapi.NewMessage(chatId, message)
	_, err := t.Bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (t *TelegramBot) SendMessages(ctx context.Context, chatIds []int64, msg string) error {
	for _, chatId := range chatIds {
		msg := tgbotapi.NewMessage(chatId, msg)
		_, err := t.Bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
