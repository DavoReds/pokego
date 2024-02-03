package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

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
	config := repl.Config{
		Next:     &map_endpoint,
		Previous: nil,
		Cache:    *pokecache.NewCache(time.Second * 5),
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		commandWords := cleanInput(scanner.Text())
		commandName := commandWords[0]
		command, ok := commandMap[commandName]

		if ok {
			err := command.Callback(&config, commandWords[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Not a command")
			continue
		}
	}
}
