package kernel

import (
	"fmt"

	commandaction "github.com/fonipts/burpium/library/action"
)

func Cmdlist() map[string]commandaction.CommandExec {
	commands := map[string]commandaction.CommandExec{
		"init":   &commandaction.StrcInit{},
		"branch": &commandaction.StrcBranch{},
		"cmd":    &commandaction.StrcCmd{},
		"help":   &StrcHelp{},
	}

	return commands
}

type StrcHelp struct {
}

func (c *StrcHelp) Aboutme() string {
	return "To see the available command"
}
func (c *StrcHelp) Execute() {
	fmt.Println("See all available command")

	commands := Cmdlist()
	for k, n := range commands {
		fmt.Println(k + " ... " + n.Aboutme())
	}
}
