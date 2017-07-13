package commands

import "strings"

type Chat struct{
	Message []string
}

func (chat Chat) Execute() string {
	return "Кто-то сказал в чате: " + strings.Join(chat.Message, " ")
}
