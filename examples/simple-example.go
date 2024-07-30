package main

import (
	"fmt"

	mini "github.com/mvstermind/mini-cli"
)

// example
func main() {

	// argument can be created by filling a Arg struct or by using NewArg function
	change := mini.Arg{
		ShortCmd: "c",
		LongCmd:  "change",
		Usage:    "Example usage",
		Required: true,
	}

	del := mini.NewArg("d", "delete", "Deletes stuff", true)
	undo := mini.NewArg("u", "undo", "Undo stuff", true)
	revert := mini.NewArg("r", "revert", "Reverts stuff", false)

	// if arugment was created using struct, it must be passed to AddArguments using "&"
	cmds := mini.AddArguments(&change, del, undo, revert)

	// return values of arguments that were found in os.Args
	argValues := cmds.Execute()

	// To acces value of flag, index into map argValues with short version of flag name
	// to acces value of "revert" value use argValues["-r"]
	for v := range argValues {
		fmt.Println(argValues[v])
	}

	//note self:
	// noticed that os.Args are str by default, find a way to make convert them afterwards
	// fmt.Printf("%t\n", argValues["-u"])
}
