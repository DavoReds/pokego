package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DavoReds/pokego/internal/pokeapi/responses"
	"github.com/DavoReds/pokego/internal/pokecache"
	"github.com/DavoReds/pokego/internal/repl"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := repl.GetCommands()
	map_endpoint := "https://pokeapi.co/api/v2/location-area"
	config := repl.State{
		Next:     &map_endpoint,
		Previous: nil,
		Cache:    *pokecache.NewCache(time.Second * 5),
		Pokedex:  map[string]responses.Pokemon{},
	}

	for {
		fmt.Print("Pokédex > ")

		if !scanner.Scan() {
			err := scanner.Err()
			if err == nil {
				break
			} else {
				fmt.Println("\nError reading input:", err)
				continue
			}
		}

		commandWords := cleanInput(scanner.Text())

		if len(commandWords) == 0 {
			continue
		}

		commandName := commandWords[0]
		command, ok := commandMap[commandName]

		if ok {
			err := command.Callback(&config, commandWords[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Not a command. Use `help` to find out what commands you can use")
			continue
		}
	}
}
