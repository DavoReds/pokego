package repl

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/DavoReds/pokego/internal/pokeapi"
	"github.com/DavoReds/pokego/internal/pokeapi/responses"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(state *State, args []string) error
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
			Description: "Exits the Pokédex",
			Callback:    exitCommand,
		},
		"map": {
			Name:        "map",
			Description: "Returns information about the next 20 locations",
			Callback:    mapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Returns information about the previous 20 locations",
			Callback:    mapbCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore an area for Pokémon",
			Callback:    exploreCommand,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokémon you just found",
			Callback:    catchCommand,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Get information about the Pokémon you've caught",
			Callback:    inspectCommand,
		},
	}

	return commands
}

func helpCommand(state *State, args []string) error {
	fmt.Print("\n")
	fmt.Println("Welcome to the Pokédex!\nUsage:")
	fmt.Print("\n")

	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}

	fmt.Print("\n")

	return nil
}

func exitCommand(state *State, args []string) error {
	os.Exit(0)
	return nil
}

func mapCommand(state *State, args []string) error {
	if state.Next == nil {
		return errors.New("No more areas. You're done!")
	}

	var response []byte

	if cached, exists := state.Cache.Get(*state.Next); exists {
		response = cached
	} else {
		apiRespose, err := pokeapi.Get(*state.Next)
		if err != nil {
			return err
		}

		response = apiRespose
	}

	var areas responses.Map
	if err := pokeapi.Parse(response, &areas); err != nil {
		return err
	}

	state.Previous = state.Next
	state.Next = areas.Next

	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func mapbCommand(state *State, args []string) error {
	if state.Previous == nil {
		return errors.New("Hold on! You haven't even started")
	}

	var response []byte

	if cached, exists := state.Cache.Get(*state.Previous); exists {
		response = cached
	} else {
		apiRespose, err := pokeapi.Get(*state.Previous)
		if err != nil {
			return err
		}
		response = apiRespose
	}

	var areas responses.Map
	if err := pokeapi.Parse(response, &areas); err != nil {
		return err
	}

	state.Next = state.Previous
	state.Previous = areas.Previous

	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func exploreCommand(state *State, args []string) error {
	if len(args) == 0 {
		return errors.New("What area do you want to explore?")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[0])
	var response []byte

	if cached, exists := state.Cache.Get(url); exists {
		response = cached
	} else {
		apiRespose, err := pokeapi.Get(url)
		if err != nil {
			return err
		}
		response = apiRespose
	}

	var area responses.Area
	if err := pokeapi.Parse(response, &area); err != nil {
		return errors.New("Something's fishy about that area")
	}

	for _, pokemon := range area.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}

	return nil
}

const catchChance = 700

func catchCommand(state *State, args []string) error {
	if len(args) == 0 {
		return errors.New("What Pokémon do you want to catch?")
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", args[0])
	var response []byte

	if cached, exists := state.Cache.Get(url); exists {
		response = cached
	} else {
		apiRespose, err := pokeapi.Get(url)
		if err != nil {
			return err
		}
		response = apiRespose
	}

	var pokemon responses.Pokemon
	if err := pokeapi.Parse(response, &pokemon); err != nil {
		return errors.New("I don't think that's a Pokémon...")
	}

	fmt.Printf("Throwing a Pokéball at %s...\n", pokemon.Name)

	catchOpportunity := catchChance - pokemon.BaseExperience
	catchAttempt := rand.Intn(catchChance)

	if catchAttempt < catchOpportunity {
		state.Pokedex[pokemon.Name] = pokemon
		fmt.Println(pokemon.Name, "was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	return nil
}

func inspectCommand(state *State, args []string) error {
	if len(args) == 0 {
		return errors.New("What Pokémon would you like me to tell you about?")
	}

	pokemon, exists := state.Pokedex[args[0]]
	if !exists {
		return errors.New("You haven't caught that Pokémon. Keep trying!")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}
