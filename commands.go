package main

import (
	"errors"
	"fmt"
	"github.com/DavoReds/pokego/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

func getCommands() map[string]cliCommand {
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

func helpCommand(conf *config) error {
	fmt.Print("\n")
	fmt.Println("Welcome to the Pokédex!\nUsage:")
	fmt.Print("\n")

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Print("\n")

	return nil
}

func exitCommand(conf *config) error {
	os.Exit(0)
	return nil
}

func mapCommand(conf *config) error {
	if conf.next == nil {
		return errors.New("No more areas. You're done!")
	}

	response, err := pokeapi.GetMapAreas(*conf.next)
	if err != nil {
		return err
	}

	conf.previous = conf.next
	conf.next = response.Next

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func mapbCommand(conf *config) error {
	if conf.previous == nil {
		return errors.New("Hold on! You haven't even started")
	}

	response, err := pokeapi.GetMapAreas(*conf.previous)
	if err != nil {
		return err
	}

	conf.next = conf.previous
	conf.previous = response.Previous

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	return nil
}
