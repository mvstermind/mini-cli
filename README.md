# mini-cli
Minimalistic library for creating CLI apps. Written in Go.

## Getting Started
To get started with `mini-cli`, install the package using `go get`.

```bash
go get github.com/mvstermind/mini-cli
```

## Usage

### Simple Usage
Here’s a basic example to get you started with `mini-cli`.

```go
package main

import (
	"fmt"

	mini "github.com/mvstermind/mini-cli"
)

func main() {
	// Create arguments
	change := mini.Arg{
		ShortCmd: "c",
		LongCmd:  "change",
		Usage:    "Example usage",
		Required: true,
	}

	del := mini.NewArg("c", "delete", "Deletes stuff", true)
	undo := mini.NewArg("u", "undo", "Undo stuff", true)
	revert := mini.NewArg("r", "revert", "Reverts stuff", false)

	// Add arguments to the command
	cmds := mini.AddArguments(&change, del, undo, revert)

	// Execute and get argument values
	argValues := cmds.Execute()

	// Access the value of the "revert" argument
	fmt.Println(argValues["-r"])
}
```

### Features
- **Minimalistic Design:** Focuses on simplicity and ease of use.
- **Flexible Argument Handling:** Easily define and handle command-line arguments.
- **Auto-generated Command List:** Automatically generates a list of required and available commands when required flags are not passed.


### Auto-generated Command List
If required flags are not passed, `mini-cli` automatically generates a list of required and available commands.

#### Example Output
```
Required value that weren't included:
-c | --change
-d | --delete
-u | --undo

List of available commands:

-c  --change
    Example usage
    Required: true

-d  --delete
    Deletes stuff
    Required: true

-u  --undo
    Undo stuff
    Required: true

-r  --revert
    Reverts stuff
    Required: false
```

## Examples
For more examples, check out the [examples section](https://github.com/mvstermind/mini-cli/tree/main/examples).

## Contributing
We welcome contributions! If you would like to contribute, please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License.
