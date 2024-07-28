package minicli

import (
	"fmt"
	"os"
)

// wille execute all of given functions (only when i implement adding fucntionality)
func (c Commands) Execute() error {
	returnval, err := c.scanInput(os.Args[1:])
	if err != nil {
		return err
	}
	// for debugging
	fmt.Println(returnval)

	return nil
}

// returns a value of argument that was seen next to ShortCmd or LongCmd
func (c Commands) scanInput(args []string) (map[string]string, error) {

	sysValues := make(map[string]string)

	c.checkIfHelp(args)

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

// checkIfHelp function scans os arguments.
// Checks if it should display help info to the user
func (c Commands) checkIfHelp(cmdArgs []string) {

	if len(cmdArgs) == 0 {
		fmt.Println("Use -h or --help")
	}

	// Iterate over all commands in structs
	for _, v := range c {

		// Iterate over all sys args
		for i := 0; i < len(cmdArgs); i++ {
			if cmdArgs[i] == v.ShortCmd || cmdArgs[i] == v.LongCmd {
				if i+1 >= len(cmdArgs) {
					fmt.Printf("\nUsage: %v\n\n", v.Usage)
				}
			}

			if cmdArgs[i] == "-h" || cmdArgs[i] == "--help" {

				fmt.Println()
				fmt.Printf("%v  %v\n", v.ShortCmd, v.LongCmd)
				fmt.Printf("\t%v\n", v.Usage)
				fmt.Printf("\tRequired: %v\n", v.Required)
			}

		}
	}

}
