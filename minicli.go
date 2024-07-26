package minicli

import (
	"fmt"
	"os"
)

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
