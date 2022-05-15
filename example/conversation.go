package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wintermonth2298/tgbotapi-conversation-extension/conv"
)

var (
	stateHandler1 = conv.State(0)
	stateHandler2 = conv.State(1)
	stateHandler3 = conv.State(2)
)

func (b *Bot) GetExampleConversationHandler() *conv.Handler {
	handler := conv.NewHandler(
		"/start",
		conv.States{
			stateHandler1: b.exampleHandlerFunc1,
			stateHandler2: b.exampleHandlerFunc2,
			stateHandler3: b.exampleHandlerFunc3,
		})

	return handler
}

func (b *Bot) exampleHandlerFunc1(ctx conv.Context, update tgbotapi.Update) conv.State {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are in handler 1")
	b.api.Send(msg)

	return stateHandler2
}

func (b *Bot) exampleHandlerFunc2(ctx conv.Context, update tgbotapi.Update) conv.State {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are in handler 2")
	b.api.Send(msg)

	return stateHandler3
}

func (b *Bot) exampleHandlerFunc3(ctx conv.Context, update tgbotapi.Update) conv.State {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are in handler 3, Finishing")
	b.api.Send(msg)

	return conv.StateCloseConversation
}
