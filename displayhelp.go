package minicli

import "fmt"

// displayHelp function scans os arguments.
// Checks if it should display help info
// if it will display help info, function will return true
func (c Commands) checkIfHelp(cmdArgs []string) bool {

	if len(cmdArgs) == 0 {
		fmt.Println("Use -h or --help")
		return true
	}

	// Iterate over all commands in structs
	for _, v := range c {

		// Iterate over all sys args
		for i := 0; i < len(cmdArgs); i++ {
			if cmdArgs[i] == v.ShortCmd || cmdArgs[i] == v.LongCmd {
				if i+1 >= len(cmdArgs) {
					fmt.Printf("\nUsage: %v\n\n", v.Usage)
					return true
				}
			}

		}
	}

	return false

}
