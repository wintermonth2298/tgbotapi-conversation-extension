package conv

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	HandlerFunc func(ctx Context, update tgbotapi.Update) State
	State       int64
	States      map[State]HandlerFunc
)

const (
	StateCloseConversation = State(-1)
)

type Handler struct {
	states       States
	users        map[int64]*conversation
	entryMessage string
}

func NewHandler(entyMessage string, states States) *Handler {
	return &Handler{
		states:       states,
		users:        make(map[int64]*conversation, 0),
		entryMessage: entyMessage,
	}
}

type conversation struct {
	ctx    Context
	states States
	state  State
}

func newConversation(states States) *conversation {
	return &conversation{
		ctx:    map[string]any{},
		states: states,
		state:  0,
	}
}

func (c *Handler) Handle(update tgbotapi.Update) bool {

	var conv *conversation
	from := update.Message.From.ID

	// create and save new conversation if entryMessage was sent
	if update.Message.Text == c.entryMessage {
		conv = newConversation(c.states)
		c.users[from] = conv
	}

	// get conversation if it exists
	for user, conversation := range c.users {
		if user == from {
			conv = conversation
			break
		}
	}

	// return if there is no conversation
	if conv == nil {
		return true
	}

	// execute handlerfunc and change state
	for state, handler := range c.states {
		if state == conv.state {
			conv.state = handler(conv.ctx, update)
			break
		}
	}

	// remove closed conversation
	if conv.state == StateCloseConversation {
		delete(c.users, from)
	}

	return false
}
