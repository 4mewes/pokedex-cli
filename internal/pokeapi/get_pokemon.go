package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/4mewes/pokedex/internal/pokecache"
)

func GetPokemonInfo(url string, cache *pokecache.Cache) (PokemonInfo, error) {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error requesting: %w", err)
			return PokemonInfo{}, fmt.Errorf("Error requesting: pokeapi.co/api/v2/pokemon/: %w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading body: %w", err)
			return PokemonInfo{}, fmt.Errorf("Error reading body: %w", err)
		}
		cache.Add(url, body)
	}

	var pokemonInfoRes PokemonInfo
	err := json.Unmarshal(body, &pokemonInfoRes)
	if err != nil {
		fmt.Println("Errr unmarshaling: %w", err)
		return PokemonInfo{}, fmt.Errorf("Error unmarshalling: %w", err)
	}
	return pokemonInfoRes, nil
}
