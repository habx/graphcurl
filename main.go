package main

import (
	"github.com/habx/graphcurl/commands"
	"github.com/habx/graphcurl/flags"
)

var version string

func main() {
	flags.Version = version
	err := commands.RootCommand.Execute()
	if err != nil {
		panic(err)
	}
}
