package commands

type Who struct{}

func (command Who) Execute() string {
	return "В этом мире нет никого лучше тебя."
}
