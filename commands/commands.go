package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]CliCommand {
	commands := map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    helpCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokédex",
			Callback:    exitCommand,
		},
		"map": {
			Name:        "map",
			Description: "Returns information about the next 20 locations",
			Callback:    mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Returns information about the next 20 locations",
			Callback:    mapbCommand,
		},
	}

	return commands
}

func helpCommand() error {
	fmt.Print("\n")
	fmt.Println("Welcome to the Pokédex!\nUsage:")
	fmt.Print("\n")

	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
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
