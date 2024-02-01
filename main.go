package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := getCommands()
	map_endpoint := "https://pokeapi.co/api/v2/location-area"
	config := config{
		next:     &map_endpoint,
		previous: nil,
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
			err := command.callback(&config)
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
