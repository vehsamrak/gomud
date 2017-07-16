package commands

import (
	"net"
	"errors"
	"github.com/Vehsamrak/gomud/console"
)

type Quit struct{
	connectionPointer *net.Conn
}

func (command Quit) GetNames() []string {
	return []string{"quit", "конец"}
}

func (command Quit) Execute() (string, error) {
	console.Client(command.connectionPointer, "Увидимся в следующий раз!")
	connection := *command.connectionPointer
	defer connection.Close()

	return "", errors.New("quit")
}
