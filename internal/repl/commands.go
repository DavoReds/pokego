package repl

import (
	"errors"
	"fmt"
	"github.com/DavoReds/pokego/internal/pokeapi"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
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

func helpCommand(conf *Config) error {
	fmt.Print("\n")
	fmt.Println("Welcome to the Pokédex!\nUsage:")
	fmt.Print("\n")

	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Print("\n")

	return nil
}

func exitCommand(conf *Config) error {
	os.Exit(0)
	return nil
}

func mapCommand(conf *Config) error {
	if conf.Next == nil {
		return errors.New("No more areas. You're done!")
	}

	var response []byte

	if cached, exists := conf.Cache.Get(*conf.Next); exists {
		response = cached
	} else {
		apiRespose, err := pokeapi.GetMapAreas(*conf.Next)
		if err != nil {
			return err
		}

		response = apiRespose
	}

	areas, err := pokeapi.ParseMapAreas(response)
	if err != nil {
		return err
	}

	conf.Previous = conf.Next
	conf.Next = areas.Next

	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func mapbCommand(conf *Config) error {
	if conf.Previous == nil {
		return errors.New("Hold on! You haven't even started")
	}

	var response []byte

	if cached, exists := conf.Cache.Get(*conf.Previous); exists {
		response = cached
	} else {
		apiRespose, err := pokeapi.GetMapAreas(*conf.Previous)
		if err != nil {
			return err
		}
		response = apiRespose
	}

	areas, err := pokeapi.ParseMapAreas(response)
	if err != nil {
		return err
	}

	conf.Next = conf.Previous
	conf.Previous = areas.Previous

	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}

	return nil
}
