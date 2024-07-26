package minicli

import (
	"fmt"
)

// wille execute all of given functions (only when i implement adding fucntionality)
func (c Commands) Execute(arg []string) error {
	returnval := c.scanInput(arg[1:])
	fmt.Println(returnval)

	return nil
}

// returns a value of argument that was seen next to ShortCmd or LongCmd
func (c Commands) scanInput(args []string) map[string]string {

	sysValues := make(map[string]string)

	for _, v := range c {
		for i := 0; i < len(args); i++ {
			if args[i] == v.ShortCmd || args[i] == v.LongCmd {
				sysValues[args[i]] = args[i+1]
			}

		}
	}

	return sysValues

}
