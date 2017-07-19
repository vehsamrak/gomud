package commands

import (
	"github.com/Vehsamrak/gomud/console"
	"net"
	"fmt"
)

type LoginCommander struct {
	ConnectionPointer *net.Conn
	ConnectionPool map[string]*net.Conn
	LoginStage string
	commandParameters []string
}

func (commander *LoginCommander) ExecuteCommand(rawCommand string) (commandResult CommandResult) {
	if rawCommand == "" {
		return
	}

	console.Server(fmt.Sprintf("[%v] Command received: %v", commander.ConnectionPointer, rawCommand))

	switch commander.LoginStage {
	case "initial":
		console.Client(commander.ConnectionPointer, "Введите свой пароль: ")
		commander.LoginStage = "1"
	case "1":
		console.Client(commander.ConnectionPointer, "Выберите расу: ")
		commander.LoginStage = "2"
	case "2":
		console.Client(commander.ConnectionPointer, "Выберите класс: ")
		commander.LoginStage = "3"
	case "3":
		console.Client(commander.ConnectionPointer, "Выберите город: ")
		commander.LoginStage = "4"
	case "4":
		console.Client(commander.ConnectionPointer, "Добро пожаловать!")
		commandResult = CommandResult{
			&GameCommander{
				commander.ConnectionPointer,
				commander.ConnectionPool,
				commander.commandParameters,
			},
		}
	}

	return commandResult
}
