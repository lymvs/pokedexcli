package main

import (
	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

var Pokedex = make(map[string]pokeapi.Pokemon)

func main() {
	startRepl()
}
