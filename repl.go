package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UsagePrinter(cmdMap *map[string]CliCommand) {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	fmt.Println()
	fmt.Println()

	for _, value := range *cmdMap {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
}

func getCommands() map[string]CliCommand {
	var cmdMap = map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Display all the location areas - 20 location areas each",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Display location areas - Go back",
			callback:    commandMapb,
		},

		"clear": {
			name:        "clear",
			description: "Clear everything in the console",
			callback:    commandClear,
		},
	}

	return cmdMap
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

func startRepl(cfg *config) {
	availableCommands := getCommands()
	UsagePrinter(&availableCommands)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, ok := availableCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
