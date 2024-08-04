package main

import (
	"fmt"
	"os"

	kernel "github.com/fonipts/burpium/library/kernel"
)

func main() {
	data := new(kernel.Cmd)

	args := os.Args
	if len(args) <= 1 {
		fmt.Println("No command found, please check by running `burpium help` to see all available command")
		os.Exit(0)
	}
	action_type := os.Args[1]
	data.Name = action_type
	data.Execute()

}
