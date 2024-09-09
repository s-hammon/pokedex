package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		reader.Scan()

		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}
		commandName := text[0]
		if command, ok := getCommands()[commandName]; ok {
			if err := command.callback(); err != nil {
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
	callback    func() error
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
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}
