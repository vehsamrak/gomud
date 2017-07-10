package commands

import "fmt"

type Command struct{}

func (t *Command) Execute() {
	fmt.Println("executed")
}
