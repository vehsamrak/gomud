package commands

import (
	"strings"
	//"fmt"
	"net"
)

type Chat struct{
	Message []string
	ConnectionPool  map[string]*net.Conn
}

func (chat Chat) Execute() string {
	//connectionPool := [...]net.Conn {}
	//for key, value := range []{chat.ConnectionPool...} {
	//	fmt.Println("Key:", key, "Value:", value)
	//}

	return "Кто-то сказал в чате: " + strings.Join(chat.Message, " ")
}
