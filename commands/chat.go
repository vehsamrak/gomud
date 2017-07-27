package commands

import (
	"strings"
	"errors"
	"github.com/Vehsamrak/gomud/player"
	"fmt"
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
	user := command.ConnectionPool[fmt.Sprint(command.sender.ConnectionPointer)]

	command.sender.toAllClients(command.ConnectionPool, fmt.Sprintf("%v говорит: \x1b[32;1m%v\x1b[0m", user.Name, message))

	return "", errors.New("Already sent")
}
