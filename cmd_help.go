package main

import "fmt"

func cmdHelp() error {
	fmt.Printf(
		"Welcome to %v! These are the available commands:\n\n",
		proompt,
	)
	for _, cmd := range getCommands() {
		fmt.Printf("%v\t%v\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
