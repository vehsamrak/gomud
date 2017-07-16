package commands

type Commandable interface {
	Namable
	Executable
}
