package main

import (
	"fmt"
	"os"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

func commandExit(_ *pokeapi.Paginate, _ []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
