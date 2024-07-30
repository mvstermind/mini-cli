package mini

import (
	"fmt"
	"os"
)

// wille execute all of given functions (only when i implement adding fucntionality)
func (c Commands) Execute() map[string]string {
	returnval := c.scanInput(os.Args[1:])

	return returnval
}

// returns a value of argument that was seen next to ShortCmd or LongCmd
func (c Commands) scanInput(args []string) map[string]string {

	sysValues := make(map[string]string)

	c.checkIfHelp(args)

	// Iterate over all commands in structs
	for _, v := range c {
		// Iterate over all sys args
		for i := 0; i < len(args); i++ {
			if args[i] == "-h" || args[i] == "--help" {
				return nil
			}

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
	if values, ok := c.checkForRequired(sysValues); ok {
		if len(values) == 1 {
			fmt.Printf("Required value that wasn't included:\n")
			fmt.Printf("%v\n", values[0])
		}
		if len(values) > 1 {
			fmt.Printf("Required value that weren't included:\n")
			for _, v := range values {
				fmt.Printf("%v\n", v)
			}
		}
		fmt.Printf("\nList of avilable commands:\n")
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
			if cmdArgs[i] == "-h" || cmdArgs[i] == "--help" {

				fmt.Println()
				fmt.Printf("%v  %v\n", v.ShortCmd, v.LongCmd)
				fmt.Printf("\t%v\n", v.Usage)
				fmt.Printf("\tRequired: %v\n", v.Required)
			}
		}
	}
	fmt.Println()

}

// checkForRequired are type of string at this point
func (c Commands) checkForRequired(foundArgs map[string]string) ([]string, bool) {

	var requiredButNotFound []string
	for _, v := range c {
		if v.Required && foundArgs[v.ShortCmd] == "" {
			requiredButNotFound = append(requiredButNotFound, fmt.Sprintf("%v | %v", v.ShortCmd, v.LongCmd))
		}
	}

	if len(requiredButNotFound) == 0 {
		return nil, false
	}
	return requiredButNotFound, true

}

func (c Commands) displayShortHelp() {

	for _, v := range c {
		fmt.Println()
		fmt.Printf("%v  %v\n", v.ShortCmd, v.LongCmd)
		fmt.Printf("\t%v\n", v.Usage)
		fmt.Printf("\tRequired: %v\n", v.Required)
	}
}
