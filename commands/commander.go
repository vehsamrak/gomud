package commands

import (
	"net"
	"strings"
	"github.com/Vehsamrak/gomud/console"
)

func ExecuteCommand(fullCommand string, connectionPointer *net.Conn, connectionPool map[string]*net.Conn) {
	if fullCommand == "" {
		return
	}

	console.Server("Command received: " + fullCommand)

	commandWithParameters := strings.Fields(fullCommand)
	commandName := commandWithParameters[0]
	commandParameters := commandWithParameters[1:]

	var command Executable

	switch commandName {
	case "test":
		fallthrough
	case "тест":
		command = Test{}

	case "who":
		fallthrough
	case "кто":
		command = Who{}

	case "look" :
		fallthrough
	case "смотреть":
		command = Look{}

	case "chat" :
		fallthrough
	case "чат":
		command = Chat{commandParameters, connectionPool}

	case "quit":
		fallthrough
	case "exit":
		fallthrough
	case "конец":
		command = Exit{}
		connection := *connectionPointer
		defer connection.Close()

	default:
		console.Client(connectionPointer, "Command not found.")

		return
	}

	commandResult, error := command.Execute()

	if error == nil {
		console.Client(connectionPointer, commandResult)
	}
}
