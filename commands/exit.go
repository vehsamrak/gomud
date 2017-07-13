package commands

type Exit struct{}

func (command Exit) Execute() (string, error) {
	return "Увидимся в следующий раз!", nil
}
