package minicli

// Arg struct containts all information about command line argument
type Arg struct {
	ShortCmd string
	LongCmd  string
	Desc     string
	Required bool
}

// type Commands keeps all of the command line arguments
// in slice of *Arg
type Commands []*Arg

// NewArg creates new argument for cli
func NewArg(short, long, desc string, required bool) *Arg {
	return &Arg{
		ShortCmd: short,
		LongCmd:  long,
		Desc:     desc,
		Required: required,
	}
}

// AddArguments function takes variadic amount of args.
// Appends them to type Commands which underlying type is
// []*Arg
func AddArguments(args ...*Arg) Commands {
	cmd := Commands{}
	cmd = append(cmd, args...)
	cmd.addPrefixToArgs()
	return cmd
}
