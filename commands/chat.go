package commands

import (
	"strings"
	"errors"
	"github.com/Vehsamrak/gomud/player"
)

type Chat struct{
	sender Sender
	ConnectionPool map[string]*player.Player
	Message []string
}

func (command *Chat) GetNames() []string {
	return []string{"chat", "чат"}
}

func (command *Chat) Execute() (string, error) {
	message := strings.Join(command.Message, " ")
	command.sender.toAllClients(command.ConnectionPool, message)

	return "", errors.New("Already sent")
}
