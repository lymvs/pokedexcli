package pokeapi

import "github.com/lymvs/pokedexcli/internal/pokecache"

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

type Paginate struct {
	Next        string
	Previous    string
	CacheResult *pokecache.Cache
}
