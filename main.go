package main

import "fmt"

type Arg struct {
	ShortCmd string
	LongCmd  string
	Desc     string
	Required bool
}

type Commands []*Arg

func NewArg(short, long, desc string, required bool) *Arg {
	return &Arg{
		ShortCmd: short,
		LongCmd:  long,
		Desc:     desc,
		Required: required,
	}
}

func AddArguments(args ...*Arg) Commands {
	cmd := Commands{}
	cmd = append(cmd, args...)
	cmd.addPrefixToArgs()
	return cmd
}

func main() {
	arg1 := NewArg("d", "delete", "deletes stuff", true)
	arg2 := NewArg("u", "undo", "undo stuff", true)
	cmds := AddArguments(arg1, arg2)

	for _, cmd := range cmds {
		fmt.Printf("ShortCmd: %s, LongCmd: %s, Desc: %s, Required: %t\n",
			cmd.ShortCmd, cmd.LongCmd, cmd.Desc, cmd.Required)
	}
}

