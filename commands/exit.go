package commands

import (
	"errors"
)

type Quit struct{
	sender Sender
}

func (command Quit) GetNames() []string {
	return []string{"quit", "конец"}
}

func (command Quit) Execute() (string, error) {
	command.sender.toClient("Увидимся в следующий раз!")
	connection := *command.sender.ConnectionPointer
	defer connection.Close()

	return "", errors.New("quit")
}
