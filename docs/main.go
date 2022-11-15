package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	"github.com/habx/graphcurl/commands"
)

func main() {
	err := doc.GenMarkdownTree(commands.RootCommand, "./")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("doc successfully generated")
	}
}
