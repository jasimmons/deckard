package main

import (
	"log"

	"github.com/jasimmons/deckard/cmd/server/command"
)

func main() {
	root := command.New()
	err := root.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
