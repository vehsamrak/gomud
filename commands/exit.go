package commands

type Exit struct{}

func (command Exit) Execute() string {
	return "Увидимся в следующий раз!"
}
