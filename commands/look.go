package commands

type Look struct{}

func (command Look) Execute() string {
	return "Ты видишь контуры этого мира."
}
