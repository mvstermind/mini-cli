package main

import (
	"fmt"
	"os"

	minicli "github.com/mvstermind/mini-cli"
)

// example
func main() {
	arg1 := minicli.NewArg("d", "delete", "deletes stuff", true)
	arg2 := minicli.NewArg("u", "undo", "undo stuff", true)
	cmds := minicli.AddArguments(arg1, arg2)

	// for debugging :3
	for _, cmd := range cmds {
		fmt.Printf("ShortCmd: %s, LongCmd: %s, Desc: %s, Required: %t\n",
			cmd.ShortCmd, cmd.LongCmd, cmd.Desc, cmd.Required)
	}
	cmds.Execute(os.Args)

}
