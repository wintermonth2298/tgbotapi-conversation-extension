package main

import (
	"flag"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(api *tgbotapi.BotAPI) *Bot {
	return &Bot{api: api}
}

func main() {
	var token string
	flag.StringVar(&token, "token", "", "telegram token")
	flag.Parse()

	api, _ := tgbotapi.NewBotAPI(token)
	bot := NewBot(api)

	convHandler := bot.GetExampleConversationHandler()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.api.GetUpdatesChan(u)

	for update := range updates {
		finished := convHandler.Handle(update)
		if !finished {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "type /start to start conversation")
		bot.api.Send(msg)
	}
}
