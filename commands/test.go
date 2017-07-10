package commands

type Test struct{}

func (command Test) Execute() string {
	return "Тест прошел."
}
