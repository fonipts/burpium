package commandaction

import (
	"fmt"
)

type CommandExec interface {
	Execute()
	Aboutme() string
}

type StrcInit struct {
}

func (c *StrcInit) Aboutme() string {
	return "Initialize the project"
}
func (c *StrcInit) Execute() {

	Initialize()
}

type StrcCmd struct {
}

func (c *StrcCmd) Aboutme() string {
	return "Execute command"
}
func (c *StrcCmd) Execute() {

	fmt.Println("This feature is working in progress")
}

type StrcBranch struct {
}

func (c *StrcBranch) Aboutme() string {
	return "Checkout to new backup"
}
func (c *StrcBranch) Execute() {

	Branching()
}
