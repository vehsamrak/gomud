package commands

import (
	"net"
	"strings"
	"github.com/Vehsamrak/gomud/console"
)

type Commander struct {
	ConnectionPointer *net.Conn
	ConnectionPool map[string]*net.Conn
	commandParameters []string
}

func (commander *Commander) ExecuteCommand(fullCommand string) {
	if fullCommand == "" {
		return
	}

	console.Server("Command received: " + fullCommand)

	commandWithParameters := strings.Fields(fullCommand)
	commandName := commandWithParameters[0]
	commander.commandParameters = commandWithParameters[1:]
	command := commander.findCommandByName(commandName)

	if command == nil {
		console.Client(commander.ConnectionPointer, "Command not found.")

		return
	}

	commandResult, error := command.Execute()

	if error == nil {
		console.Client(commander.ConnectionPointer, commandResult)
	}
}

func (commander *Commander) findCommandByName(requestedCommandName string) Commandable {
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
func (commander *Commander) createAllCommands() []Commandable  {
	return []Commandable{
		Chat{commander.commandParameters, commander.ConnectionPool},
		Quit{commander.ConnectionPointer},
		Look{},
		Test{},
		Who{},
	}
}
