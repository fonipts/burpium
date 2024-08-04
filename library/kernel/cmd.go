package kernel

import (
	"fmt"
)

type Cmd struct {
	Name string
}

func (c *Cmd) Execute() {

	commands := Cmdlist()

	if command := commands[c.Name]; command == nil {
		fmt.Println("No such command found, please run `burpium help` to see all available command")
	} else {
		command.Execute()
	}

}
