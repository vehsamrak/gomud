package commands

type Commandable interface {
	Execute() (string, error)
	GetNames() []string
}
