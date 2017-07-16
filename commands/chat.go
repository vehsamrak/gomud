package commands

import (
	"strings"
	"net"
	"github.com/Vehsamrak/gomud/console"
	"errors"
)

type Chat struct{
	Message []string
	ConnectionPool  map[string]*net.Conn
}

func (chat Chat) GetNames() []string {
	return []string{"chat", "чат"}
}

func (chat Chat) Execute() (string, error) {
	message := strings.Join(chat.Message, " ")

	for _, connection := range chat.ConnectionPool {
		console.Client(connection, "Кто-то сказал в чате: " + message)
	}

	return "", errors.New("Already sent")
}
