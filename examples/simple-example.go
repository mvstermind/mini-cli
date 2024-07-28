package main

import (
	"fmt"

	minicli "github.com/mvstermind/mini-cli"
)

// example
func main() {
	del := minicli.NewArg("d", "delete", "Deletes stuff", true, "rm")
	undo := minicli.NewArg("u", "undo", "Undo stuff", true, 2)
	revert := minicli.NewArg("r", "revert", "Reverts stuff", false, "-2")

	cmds := minicli.AddArguments(del, undo, revert)

	argValues := cmds.Execute()

	// To acces value of flag, index into map argValues with shorter version of flag name
	fmt.Println(argValues["-u"])

	// noticed that os.Args are str by default, find a way to make convert them afterwards
	// fmt.Printf("%t\n", argValues["-u"])
}
