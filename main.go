package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/williammartin/nsjam/commands"
)

type command struct {
	name        string
	description string
	command     interface{}
}

func main() {
	parser := flags.NewParser(nil, flags.HelpFlag|flags.PassDoubleDash)

	commands := []command{
		{"version", "returns the semantic version in use", &commands.Version{}},
		{"list-namespaces", "lists the namespaces for a process", &commands.ListNamespaces{}},
	}

	for _, command := range commands {
		_, err := parser.AddCommand(
			command.name,
			command.description,
			"",
			command.command,
		)

		if err != nil {
			panic(err)
		}
	}

	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
