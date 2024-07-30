package main

import (
	"fmt"
	"net/http"

	"github.com/mvstermind/mini-cli"
)

func request() {

	arg := mini.Arg{
		ShortCmd: "u",
		LongCmd:  "url",
		Usage:    "pass in url to get body",
		Required: true,
	}

	commands := mini.AddArguments(&arg)
	exec := commands.Execute()

	resp, err := http.Get(exec["-u"])
	if err != nil {
		return
	}

	body := resp.StatusCode
	fmt.Println(body)

}
