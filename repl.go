package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
	"github.com/lymvs/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Paginate) error
}

func cleanInput(text string) []string {
	var words []string

	words = strings.Fields(strings.ToLower(text))

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

func startRepl() {
	p := &pokeapi.Paginate{
		CacheResult: pokecache.NewCache(5 * time.Second),
	}
	commands := getCommands()
	prompt := "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if command, ok := commands[input[0]]; ok {
			err := command.callback(p)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Print("Unkown command")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
