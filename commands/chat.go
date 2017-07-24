package commands

import (
	"strings"
	"errors"
	"github.com/Vehsamrak/gomud/console"
	"github.com/Vehsamrak/gomud/player"
)

type Chat struct{
	sender Sender
	ConnectionPool map[string]*player.Player
	Message []string
}

func (chat Chat) GetNames() []string {
	return []string{"chat", "чат"}
}

func (chat Chat) Execute() (string, error) {
	message := strings.Join(chat.Message, " ")

	for _, user := range chat.ConnectionPool {
		connection := user.ConnectionPointer
		console.Client(connection, "Chat: " + message)
	}

	return "", errors.New("Already sent")
}
