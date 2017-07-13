package commands

import (
	"strings"
	"net"
	"github.com/Vehsamrak/gomud/console"
)

type Chat struct{
	Message []string
	ConnectionPool  map[string]*net.Conn
}

func (chat Chat) Execute() string {
	message := strings.Join(chat.Message, " ")

	for _, connection := range chat.ConnectionPool {
		console.Client(connection, "Кто-то сказал в чате: " + message)
	}

	return "Кто-то сказал в чате: " + message
}
