package commands

import (
	"strings"
	"github.com/Vehsamrak/gomud/player"
	"fmt"
)

type GameCommander struct {
	Sender
	ConnectionPool map[string]*player.Player
	commandParameters []string
}

func (commander *GameCommander) ExecuteCommand(rawCommand string) (commandResult CommandResult) {
	if rawCommand == "" {
		return
	}

	rawCommand = commander.encodeToUtf8(rawCommand)
	user := commander.ConnectionPool[fmt.Sprint(commander.ConnectionPointer)]

	commander.Sender.toServer(
		fmt.Sprintf(
			"[%v] Command received: %v",
			user.Name,
			rawCommand,
		))

	commandWithParameters := strings.Fields(rawCommand)
	commandName := commandWithParameters[0]
	commander.commandParameters = commandWithParameters[1:]
	command := commander.findCommandByName(commandName)

	if command == nil {
		commander.Sender.toClient("Command not found.")

		return
	}

	commandOutput, error := command.Execute()

	if error == nil {
		commander.Sender.toClient(commandOutput)
	}

	return
}

func (commander *GameCommander) findCommandByName(requestedCommandName string) Commandable {
	var commandable Commandable

	for _, command := range commander.createAllCommands() {
		for _, commandName := range command.GetNames() {
			if strings.HasPrefix(commandName, requestedCommandName) {
				commandable = command
			}
		}
	}

	return commandable
}

// All game commands are created by this method
func (commander *GameCommander) createAllCommands() []Commandable  {
	return []Commandable{
		&Chat{commander.Sender, commander.ConnectionPool, commander.commandParameters},
		Quit{commander.Sender},
		Look{},
		Test{},
		Who{},
	}
}
