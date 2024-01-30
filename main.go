package main

import (
	"bufio"
	"fmt"
	"github.com/DavoReds/pokego/commands"
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
	commandMap := commands.GetCommands()

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
			err := command.Callback()
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
