# 🎉 Mini Cli - The Teeny-Tiny Library for Mighty CLI Apps! 🚀

Welcome to **Mini Cli**, the *itty-bitty, teeny-weeny* Go library that's small in size but **HUGE** in potential! It's your go-to toolkit for creating command-line apps with a minimalistic flair. Because why complicate things when you can keep it simple and snazzy? 😎

## How to Dive In 🏊‍♂️

Want to get started with **mini-cli**? Of course, you do! Just run this magical command and you're golden:

```bash
go get github.com/mvstermind/mini-cli
```

Boom! You just snagged yourself a copy of the coolest CLI library in town. 🎉

## How to Play 🎮

### Simple Usage - The Bare Essentials 🛠️

Here’s a basic example that’ll get you up and running with **Mini Cli** faster than you can say "command-line interface!"

```go
package main

import (
	"fmt"

	mini "github.com/mvstermind/mini-cli"
)

func main() {
	// Create a bunch of arguments (because one is never enough!)
	change := mini.Arg{
		ShortCmd: "c",
		LongCmd:  "change",
		Usage:    "Example usage",
		Required: true, // This one’s important!
	}

	del := mini.NewArg("c", "delete", "Deletes stuff", true)
	undo := mini.NewArg("u", "undo", "Undo stuff", true)
	revert := mini.NewArg("r", "revert", "Reverts stuff", false)

	// Toss those arguments into your command soup
	cmds := mini.AddArguments(&change, del, undo, revert)

	// Stir it all up and get those juicy argument values
	argValues := cmds.Execute()

	// Spill the beans on what "revert" is all about
	fmt.Println(argValues["-r"])
}
```

### Why You'll Love It ❤️

- **Minimalistic Magic:** Focuses on keeping things simple and easy-peasy. ✨
- **Flexible as a Rubber Band:** Easily define and handle command-line arguments like a boss.
- **Auto-generated Command List:** If you forget to pass those all-important flags, don’t sweat it! **mini-cli** has your back and will auto-generate a list of what you’re missing. 📝

### Auto-generated Command List - The Safety Net 🕸️

If you skip out on those required flags, **mini-cli** will jump in and save the day by automatically listing what you need. Here’s a sneak peek of what that looks like:

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

See? No more excuses for missing flags! **Mini Cli** has got you covered. 🎯

## Peek at Some Examples 👀

Want more examples? Of course, you do! Check out our [examples section](https://github.com/mvstermind/mini-cli/tree/main/examples) for a treasure trove of cool tricks you can do with **mini-cli**.

## Contributing - Join the Fun! 🎉

We love new ideas and improvements! If you want to jump into the **Mini Cli** madness, fork the repo, whip up some magic, and submit a pull request. Let’s make CLI apps awesome together! 💪

## License to Chill 🎵

This project is licensed under the MIT License, so feel free to kick back, relax, and enjoy using **mini-cli** without a care in the world. 🏖️
