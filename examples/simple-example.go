package main

import (
	"fmt"

	mini "github.com/mvstermind/mini-cli"
)

func main() {
	// Create arguments
	change := mini.Arg{
		ShortCmd: "c",
		LongCmd:  "change",
		Usage:    "Example usage",
		Required: true,
	}

	revert := mini.NewArg("r", "revert", "Reverts stuff", false)

	// Add arguments to the command
	cmds := mini.AddArguments(&change, revert)

	// Execute and get argument values
	argValues := cmds.Execute()

	// Access the value of the "revert" argument
	fmt.Println(argValues["-c"])
}
