package minicli

import (
	"fmt"
)

// wille execute all of given functions (only when i implement adding fucntionality)
func (c Commands) Execute(arg []string) error {
	_, err := c.scanInput(arg[1:])
	if err != nil {
		return err
	}
	// for debugging
	// fmt.Println(returnval)

	return nil
}

// returns a value of argument that was seen next to ShortCmd or LongCmd
func (c Commands) scanInput(args []string) (map[string]string, error) {

	sysValues := make(map[string]string)

	c.displayHelp(args)

	// Iterate over all commands in structs
	for _, v := range c {
		// Iterate over all sys args
		for i := 0; i < len(args); i++ {

			if args[i] == v.ShortCmd || args[i] == v.LongCmd {
				// Check if there's a next argument before accessing it
				if i+1 < len(args) && args[i+1] != "" {
					sysValues[args[i]] = args[i+1]
				}
			}
		}
	}

	return sysValues, nil
}

// displayHelp function scans os arguments.
// Checks if it should display help info
// if it will display help info, function will return true
func (c Commands) displayHelp(cmdArgs []string) bool {

	if len(cmdArgs) == 0 {
		fmt.Println("Use -h or --help")
		return true
	}

	// Iterate over all commands in structs
	for _, v := range c {

		// Iterate over all sys args
		for i := 0; i < len(cmdArgs); i++ {
			if cmdArgs[i] == v.ShortCmd || cmdArgs[i] == v.LongCmd && cmdArgs[i+1] == "" {
				fmt.Printf("\nUsage: %v\n\n", v.Usage)
				return true

			}

		}
	}

	return false

}
