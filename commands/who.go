package commands

type Who struct{}

func (command Who) GetNames() []string {
	return []string{"who", "кто"}
}

func (command Who) Execute() (string, error) {
	return "В этом мире нет никого лучше тебя.", nil
}
