package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	replStart()
}

// lowercase and split a string in words
func cleanInput(text string) []string {
	text = strings.TrimSpace(strings.ToLower(text))
	if text == "" {
		return []string{}
	}
	return strings.Fields(text)
}

// start REPL and take commands
func replStart() {
	userInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >/")
		userInput.Scan()
		words := cleanInput(userInput.Text())
		if len(words) == 0 {
			continue
		}

		// Command checking
		commandName := words[0]
		if len(words) > 1 {
		}

		command, exists := commandLookup()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// Command Stuff
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Exit
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Help
func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commandLookup() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandLookup() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
