package main

import (
	"fmt"
	"os"
	"github.com/lymvs/pokedexcli/internal/locationarea"
)

func commandExit(_ *locationarea.Paginate) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
