package commands

import (
	"net"
	"strings"
	"github.com/Vehsamrak/gomud/console"
)

type GameCommander struct {
	ConnectionPointer *net.Conn
	ConnectionPool map[string]*net.Conn
	commandParameters []string
}

func (commander *GameCommander) ExecuteCommand(rawCommand string) (commandResult CommandResult) {
	if rawCommand == "" {
		return
	}

	console.Server("Command received: " + rawCommand)

	commandWithParameters := strings.Fields(rawCommand)
	commandName := commandWithParameters[0]
	commander.commandParameters = commandWithParameters[1:]
	command := commander.findCommandByName(commandName)

	if command == nil {
		console.Client(commander.ConnectionPointer, "Command not found.")

		return
	}

	commandOutput, error := command.Execute()

	if error == nil {
		console.Client(commander.ConnectionPointer, commandOutput)
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
		Chat{commander.commandParameters, commander.ConnectionPool},
		Quit{commander.ConnectionPointer},
		Look{},
		Test{},
		Who{},
	}
}
