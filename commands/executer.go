package commands

type Executer interface {
	ExecuteCommand(rawCommand string) (commandResult CommandResult)
}
