package anthropic

import (
	"fmt"
	"strings"
)

const (
	MessageSenderSystem    = "system"
	MessageSenderHuman     = "\n\nHuman"
	MessageSenderAssistant = "\n\nAssistant"
)

type UserType string

type Message struct {
	Sender  UserType // The sender's name (e.g., "Human" or a username)
	Content string   // The content of the message
}

func (m *Message) marshal() string {
	if m.Sender == MessageSenderSystem {
		return m.Content
	}
	return fmt.Sprintf("%s: %s", m.Sender, m.Content)
}

func GetPromptFromMessages(msg []*Message) string {
	var prompt = make([]string, len(msg))
	for i, m := range msg {
		prompt[i] = m.marshal()
	}
	return strings.Join(prompt, "")
}

func GetPromptFromString(question string) string {
	return fmt.Sprintf("\n\nHuman: %s\n\nAssistant:", question)
}

func GetPromptFromStringWithSystemMessage(system, human string) string {
	return fmt.Sprintf("%s%s", system, GetPromptFromString(human))
}
