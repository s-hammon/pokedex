package main

import (
	"os"
	"os/exec"
)

func cmdClear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
