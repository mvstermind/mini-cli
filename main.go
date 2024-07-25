package minicli

import (
	"fmt"
)

type Command struct {
	ShortCmd []string
	LongCmd  []string
	Desc     string
	Required bool
}

func NewCommand(short, long string, desc string, required bool) *Command {
	return &Command{
		ShortCmd: []string{short},
		LongCmd:  []string{long},
		Desc:     desc,
		Required: required,
	}

}

func addPrefix(c *Command) *Command {

	for i := 0; i < len(c.LongCmd); i++ {
		addedVal := fmt.Sprintf("--%v", i)
		c.LongCmd[i] = addedVal

	}
	for i := 0; i < len(c.ShortCmd); i++ {
		addedVal := fmt.Sprintf("-%v", i)
		c.ShortCmd[i] = addedVal
	}
	return &Command{
		ShortCmd: c.ShortCmd,
		LongCmd:  c.LongCmd,
		Desc:     c.Desc,
		Required: c.Required,
	}

}
