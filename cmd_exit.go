package main

import (
	"fmt"
	"os"
)

func cmdExit(cfg *pokeConfig) error {
	fmt.Println("See ya!")
	os.Exit(0)
	return nil
}
