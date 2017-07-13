package commands

type Who struct{}

func (command Who) Execute() (string, error) {
	return "В этом мире нет никого лучше тебя.", nil
}
