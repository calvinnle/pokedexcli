package main

import (
    "fmt"
)

func commandHelp(cfg *config, arg string) error {
    commands := getCommands()
	fmt.Println("Help service requested - Welcome to the help menu")
    UsagePrinter(&commands)
	return nil
}
