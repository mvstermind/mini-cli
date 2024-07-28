package main

import (
	minicli "github.com/mvstermind/mini-cli"
)

// example
func main() {
	// specify new arguments
	arg1 := minicli.NewArg("d", "delete", "Deletes stuff", true)
	arg2 := minicli.NewArg("u", "undo", "Undo stuff", true)

	// merge them into slice
	cmds := minicli.AddArguments(arg1, arg2)

	// look for given agrs in os.Args
	cmds.Execute()

}
