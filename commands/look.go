package commands

type Look struct{}

func (command Look) GetNames() []string {
	return []string{"look", "смотреть"}
}

func (command Look) Execute() (string, error) {
	return "Ты видишь контуры этого мира.", nil
}
