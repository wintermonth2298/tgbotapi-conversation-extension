package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/wintermonth2298/tgbotapi-conversation-extension/conv"
)

var (
	stateSayHi       = conv.State(0)
	stateAskQuestion = conv.State(1)
	stateSayBye      = conv.State(2)
)

func (b *Bot) GetExampleConversationHandler() *conv.Handler {
	handler := conv.NewHandler(
		"/start",
		conv.States{
			stateSayHi:       b.exampleHandlerFunc1,
			stateAskQuestion: b.exampleHandlerFunc2,
			stateSayBye:      b.exampleHandlerFunc3,
		})

	return handler
}

func (b *Bot) exampleHandlerFunc1(ctx conv.Context, update tgbotapi.Update) conv.State {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are in handler 1")
	b.api.Send(msg)

	return stateAskQuestion
}

func (b *Bot) exampleHandlerFunc2(ctx conv.Context, update tgbotapi.Update) conv.State {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are in handler 2")
	b.api.Send(msg)

	return stateSayBye
}

func (b *Bot) exampleHandlerFunc3(ctx conv.Context, update tgbotapi.Update) conv.State {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are in handler 3, Finishing")
	b.api.Send(msg)

	return conv.StateCloseConversation
}
