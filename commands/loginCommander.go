package commands

import (
	"github.com/Vehsamrak/gomud/player"
	"fmt"
)

type LoginCommander struct {
	Sender
	ConnectionPool map[string]*player.Player
	LoginStage string
	commandParameters []string
}

func (commander *LoginCommander) ExecuteCommand(rawCommand string) (commandResult CommandResult) {
	if rawCommand == "" {
		return
	}

	user := commander.ConnectionPool[fmt.Sprint(commander.ConnectionPointer)]

	switch commander.LoginStage {
	case "charset":
		availableCharsets := map[string]bool{
			"1": true,
			"2": true,
			"3": true,
			"4": true,
		}

		charsetMap := map[string]string {
			"1": DEFAULT_ENCODING,
			"2": ENCODING_KOI8,
			"3": ENCODING_WINDOWS_1251,
			"4": ENCODING_WINDOWS_1252,
		}

		if availableCharsets[rawCommand] {
			commander.charset = charsetMap[rawCommand]
			user.Codepage = commander.charset
			commander.Sender.toClient("Введите ваше имя: ")
			commander.LoginStage = "name"
		} else {
			commander.charset = DEFAULT_ENCODING
			commander.Sender.toClient("BBEDUTE HOMEP KODUPOBKU: 1) UTF-8, 2) KOI8-R, 3) Windows-1251, 4) Windows-1252")
		}
	case "name":
		user.Name = rawCommand

		defer commander.Sender.toServer(fmt.Sprintf("%v logged in.", user.Name))

		commander.Sender.toClient(fmt.Sprintf("Добро пожаловать, %v!", user.Name))
		commandResult = CommandResult{
			&GameCommander{
				Sender: commander.Sender,
				ConnectionPool: commander.ConnectionPool,
				commandParameters: commander.commandParameters,
			},
		}
	}

	commander.Sender.toServer(
		fmt.Sprintf(
			"[%v] Command received: %v",
			commander.ConnectionPointer,
			commander.encodeToUtf8(rawCommand),
		))

	return commandResult
}
