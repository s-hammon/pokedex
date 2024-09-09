package main

import (
	"os"
	"os/exec"
)

func cmdClear(cfg *pokeConfig) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
