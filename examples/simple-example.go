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

	del := mini.NewArg("c", "delete", "Deletes stuff", true)
	undo := mini.NewArg("u", "undo", "Undo stuff", true)
	revert := mini.NewArg("r", "revert", "Reverts stuff", false)

	// Add arguments to the command
	cmds := mini.AddArguments(&change, del, undo, revert)

	// Execute and get argument values
	argValues := cmds.Execute()

	// Access the value of the "revert" argument
	fmt.Println(argValues["-r"])
}

