package commands

import (
	"github.com/Vehsamrak/gomud/console"
	"net"
	"fmt"
)

type LoginCommander struct {
	Sender
	ConnectionPool map[string]*net.Conn
	LoginStage string
	commandParameters []string
}

func (commander *LoginCommander) ExecuteCommand(rawCommand string) (commandResult CommandResult) {
	if rawCommand == "" {
		return
	}

	switch commander.LoginStage {
	case "charset":
		availableCharsets := map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		}

		charsetMap := map[string]string {
			"1": DEFAULT_ENCODING,
			"2": "koi8-r",
			"3": "windows-1251",
		}

		if availableCharsets[rawCommand] {
			commander.charset = charsetMap[rawCommand]
			commander.Sender.toClient("Введите ваше имя: ")
			commander.LoginStage = "name"
		} else {
			commander.charset = DEFAULT_ENCODING
			commander.Sender.toClient("BBEDUTE HOMEP KODUPOBKU: 1) UTF-8, 2) KOI8-R, 3) Windows-1251")
		}
	case "name":
		commander.Sender.toClient("Введите свой пароль: ")
		commander.LoginStage = "2"
	case "2":
		commander.Sender.toClient("Выберите расу: ")
		commander.LoginStage = "3"
	case "3":
		commander.Sender.toClient("Выберите класс: ")
		commander.LoginStage = "4"
	case "4":
		commander.Sender.toClient("Добро пожаловать!")
		commandResult = CommandResult{
			&GameCommander{
				commander.ConnectionPointer,
				commander.ConnectionPool,
				commander.commandParameters,
			},
		}
	}

	console.Server(
		fmt.Sprintf(
			"[%v] Command received: %v",
			commander.ConnectionPointer,
			commander.translateToUtf8(rawCommand),
		))

	return commandResult
}
