package mini

import (
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	commands := Commands{
		{ShortCmd: "-f", LongCmd: "--file", Usage: "Specify the file", Required: true},
		{ShortCmd: "-v", LongCmd: "--verbose", Usage: "Enable verbose mode", Required: false},
	}
	t.Run("Test case with required argument", func(t *testing.T) {
		os.Args = []string{"cmd", "-f", "testfile"}
		result := commands.Execute()
		if result["-f"] != "testfile" {
			t.Errorf("Expected %v, got %v", "testfile", result["-f"])
		}

	})

	t.Run("Test case with missing required argument", func(t *testing.T) {
		os.Args = []string{"cmd"}
		result := commands.Execute()
		if result != nil {
			t.Errorf("Expected %v, got %v", nil, result)
		}
	})
}

func TestScanInput(t *testing.T) {
	commands := Commands{
		{ShortCmd: "-f", LongCmd: "--file", Usage: "Specify the file", Required: true},
		{ShortCmd: "-v", LongCmd: "--verbose", Usage: "Enable verbose mode", Required: false},
	}

	args := []string{"-f", "testfile", "-v"}
	result := commands.scanInput(args)
	if result["-f"] != "testfile" {
		t.Errorf("Expected %v, got %v", "testfile", result["-f"])
	}
	if result["-v"] != "" {
		t.Errorf("Expected %v, got %v", "", result["-v"])
	}
}

func TestCheckIfHelp(t *testing.T) {
	commands := Commands{
		{ShortCmd: "-f", LongCmd: "--file", Usage: "Specify the file", Required: true},
		{ShortCmd: "-v", LongCmd: "--verbose", Usage: "Enable verbose mode", Required: false},
	}

	cmdArgs := []string{"-h"}
	commands.checkIfHelp(cmdArgs)
	// This function doesn't return anything, just ensure it doesn't panic
}

func TestCheckForRequired(t *testing.T) {
	commands := Commands{
		{ShortCmd: "-f", LongCmd: "--file", Usage: "Specify the file", Required: true},
		{ShortCmd: "-v", LongCmd: "--verbose", Usage: "Enable verbose mode", Required: false},
	}

	t.Run("Test case to see if required flags are recognized", func(t *testing.T) {

		foundArgs := map[string]string{"-f": "testfile"}
		missing, ok := commands.checkForRequired(foundArgs)
		if ok || len(missing) != 0 {
			t.Errorf("Expected %v, got %v", false, ok)
		}
	})

	t.Run("Test case to see if it will run without required flag", func(t *testing.T) {
		foundArgs := map[string]string{"-v": ""}
		missing, ok := commands.checkForRequired(foundArgs)
		if !ok || len(missing) != 1 {
			t.Errorf("Expected %v, got %v", true, ok)
		}
	})

}

func TestDisplayShortHelp(t *testing.T) {
	commands := Commands{
		{ShortCmd: "-f", LongCmd: "--file", Usage: "Specify the file", Required: true},
		{ShortCmd: "-v", LongCmd: "--verbose", Usage: "Enable verbose mode", Required: false},
	}

	// This function doesn't return anything, just ensure it doesn't panic
	commands.displayShortHelp()
}
