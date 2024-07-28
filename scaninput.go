package minicli

import (
	"fmt"
	"os"
)

// wille execute all of given functions (only when i implement adding fucntionality)
func (c Commands) Execute() map[string]any {
	returnval := c.scanInput(os.Args[1:])

	return returnval
}

// returns a value of argument that was seen next to ShortCmd or LongCmd
func (c Commands) scanInput(args []string) map[string]any {

	sysValues := make(map[string]any)

	c.checkIfHelp(args)

	// Iterate over all commands in structs
	for _, v := range c {
		// Iterate over all sys args
		for i := 0; i < len(args); i++ {

			if args[i] == v.ShortCmd || args[i] == v.LongCmd {
				// Check if there's a next argument before accessing it
				if i+1 < len(args) && args[i+1] != "" {
					// append to an map with only v.ShortCMD for faster dev time
					sysValues[v.ShortCmd] = args[i+1]
				}
			}

		}
	}

	// If this is true, it means required flag was ommitted.
	// program cannot work properly,
	if c.checkForRequired(sysValues) == true {
		fmt.Printf("Required value wasn't included\nList of avilable commands:\n")
		c.displayShortHelp()
		return nil
	}

	return sysValues
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
				fmt.Printf("\t\tDefault value: %v\n", v.Default)
			}
		}
	}

}

// foundArgs are type of string at this point
func (c Commands) checkForRequired(foundArgs map[string]any) bool {

	for _, v := range c {

		if v.Required == true && foundArgs[v.ShortCmd] == nil {
			return true
		}
	}
	return false

}

func (c Commands) displayShortHelp() {

	for _, v := range c {
		fmt.Println()
		fmt.Printf("%v  %v\n", v.ShortCmd, v.LongCmd)
		fmt.Printf("\t%v\n", v.Usage)
		fmt.Printf("\tRequired: %v\n", v.Required)
		fmt.Printf("\t\tDefault value: %v\n", v.Default)
	}
}
