package main

type Arg struct {
	ShortCmd string
	LongCmd  string
	Desc     string
	Required bool
}

type Commands []*Arg

func NewArg(short, long string, desc string, required bool) *Arg {
	return &Arg{
		ShortCmd: short,
		LongCmd:  long,
		Desc:     desc,
		Required: required,
	}

}

func AddArguments(args ...*Arg) *Commands {
	cmd := Commands{}
	cmd = append(cmd, args...)
	cmd.addPrefixToArgs()
	return &cmd
}

func main() {

	Arg := NewArg("d", "delete", "deletes stuff", true)
	Arg2 := NewArg("u", "undo", "undo stuff", true)
	_ = AddArguments(Arg, Arg2)

}
