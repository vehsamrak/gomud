package commands

type Test struct{}

func (command Test) Execute() (string, error) {
	return "Тест прошел.", nil
}
