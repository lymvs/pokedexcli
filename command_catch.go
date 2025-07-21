package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

func commandCatch(c *pokeapi.Paginate, args []string) error {
	base_url := "https://pokeapi.co/api/v2/pokemon/"
	pokemon_name := args[0]
	url := base_url + pokemon_name

	var data []byte

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("Response failed")
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	pokemon := pokeapi.Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	rnn := rng.Float64()
	catch_probability := 50.0 / float64(pokemon.BaseExperience)
	if rnn < catch_probability {
		fmt.Printf("%s escaped!\n", pokemon_name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon_name)
		Pokedex[pokemon_name] = pokemon
	}

	return nil
}
