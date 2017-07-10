package commands

type Executable interface {
	Execute() string
}

