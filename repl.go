package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/s-hammon/pokedex/internal/pokeapi"
)

type pokeConfig struct {
	pokeClient *pokeapi.Client
	nextURL    *string
	prevURL    *string
}

func startRepl(cfg *pokeConfig) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		reader.Scan()

		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]
		var args []string
		if len(text) > 1 {
			args = text[1:]
		}

		if command, ok := getCommands()[commandName]; ok {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			handleCmd(commandName)
			continue
		}
	}
}

func handleCmd(text string) {
	if text == "" {
		return
	}
	defer handleInvalidCmd(text)
}

func handleInvalidCmd(text string) {
	if text == "" {
		return
	}
	defer printUnknownCommand(text)
}

func printUnknownCommand(text string) {
	fmt.Println(text, ": command not found")
}

func printPrompt() {
	fmt.Print(proompt, " > ")
}

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeConfig, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"clear": {
			name:        "clear",
			description: "Clear the screen",
			callback:    cmdClear,
		},
		"help": {
			name:        "help",
			description: "Display this help message",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the Pokedex",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "Display the next page of areas on map",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of areas on map",
			callback:    cmdMapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location area and see the pokemon that can be found there",
			callback:    cmdExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			callback:    cmdCatch,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}
