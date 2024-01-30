package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Returns information about the next 20 locations",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns information about the next 20 locations",
			callback:    mapbCommand,
		},
	}

	return commands
}

func helpCommand() error {
	fmt.Print("\n")
	fmt.Println("Welcome to the Pokédex!\nUsage:")
	fmt.Print("\n")

	for _, command := range commands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Print("\n")

	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}

func mapCommand() error {
	return nil
}

func mapbCommand() error {
	return nil
}
