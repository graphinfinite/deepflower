package telegram

import (
	"context"
	"deepflower/internal/model"
	"deepflower/internal/observer"
	"fmt"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog"
)

type TelegramBot struct {
	Bot       *tgbotapi.BotAPI
	log       *zerolog.Logger
	callbacks map[string]*string
}

func NewBot(token string, buffer int, client http.Client, debug bool, logger *zerolog.Logger) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(token, &client)
	bot.Debug = debug
	bot.Buffer = buffer
	var callbacks = make(map[string]*string)
	if err != nil {
		return &TelegramBot{}, err
	}
	return &TelegramBot{Bot: bot, log: logger, callbacks: callbacks}, nil
}

type CallBackPayload struct {
	TgId      int
	ProcessId string
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
	go func(outChan chan observer.Event) {
		for upd := range updateschan {

			if upd.Message != nil {
				t.log.Debug().Msgf("Update Msg [%s] %s", upd.Message.From.UserName, upd.Message.Text)
				var message string
				var msg tgbotapi.MessageConfig
				usertg := model.UserTelegram{
					TgId:           upd.Message.From.ID,
					TgChatId:       upd.Message.Chat.ID,
					TgFirstName:    upd.Message.From.FirstName,
					TgLastName:     upd.Message.From.LastName,
					TgUserName:     upd.Message.From.UserName,
					TgLanguageCode: upd.Message.From.LanguageCode,
				}

				switch upd.Message.Text {
				case "/start":
					message = "Send '/auth' for registration"
					msg = tgbotapi.NewMessage(upd.Message.Chat.ID, message)
					t.Bot.Send(msg)

				case "/auth":
					event.Topic = "bot/registration"
					event.Payload = usertg
					outChan <- event
				default:
					message = "unknow"
					msg = tgbotapi.NewMessage(upd.Message.Chat.ID, message)
					t.Bot.Send(msg)
				}
				// end message handling
			}

			if upd.CallbackQuery != nil {
				t.log.Debug().Msgf("BOT/CallbackQuery %#v ", upd.CallbackQuery)

				switch upd.CallbackQuery.Data[:1] {

				// TODO
				//pc/%s/confirm
				case "pc":
					event.Topic = "bot/pc/confirm"

					processId := strings.Split(upd.CallbackQuery.Data, "/")[1]
					fmt.Printf("BOT PC: %s", processId)

					event.Payload = CallBackPayload{TgId: upd.CallbackQuery.From.ID, ProcessId: processId}
					outChan <- event
					t.Bot.DeleteMessage(tgbotapi.NewDeleteMessage(upd.CallbackQuery.Message.Chat.ID, upd.CallbackQuery.Message.MessageID))
				}
			}

		}
	}(outChan)
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

func (t *TelegramBot) SendMessagesWithCallbacks(ctx context.Context, chatIds []int64, msg string, callback map[string]string) error {
	var buttons []tgbotapi.InlineKeyboardButton
	for key, val := range callback {
		t.callbacks[key] = &val
		buttons = append(buttons, tgbotapi.InlineKeyboardButton{Text: key, CallbackData: t.callbacks[key]})
	}

	for _, chatId := range chatIds {
		msg := tgbotapi.NewMessage(chatId, msg)

		msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{buttons}}

		_, err := t.Bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
