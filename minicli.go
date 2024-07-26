package minicli

// Arg struct containts all information about command line argument
//
// ShortCmd: Is equalivent to describing short version of command.
// Example: "u" will have to used as "-u" to make it run.
//
// LongCmd: Works the same way as ShortCmd.
// String given as argument will have "--" added to it.
// Example: "usage" will have to be used as "--usage" to make it run.
//
// Usage: Short description of what this flag does, and how to use it.
// "Usage" will be displayed when empty VALID flag is provided.
//
// Required: Field that determines if flag is necessary for program to run.
type Arg struct {
	ShortCmd string
	LongCmd  string
	Usage    string
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
		Usage:    desc,
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
