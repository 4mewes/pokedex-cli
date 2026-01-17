package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/4mewes/pokedex/internal/pokecache"
)

func GetLocationArea(url string, cache *pokecache.Cache) (LocationArea, error) {
	//check cache
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error requesting: %w", err)
			return LocationArea{}, fmt.Errorf("Error requesting: pokeapi.co/api/v2/location-area/: %w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading body: %w", err)
			return LocationArea{}, fmt.Errorf("Error reading body: %w", err)
		}
		cache.Add(url, body)
	}

	var locationAreaRes LocationArea
	err := json.Unmarshal(body, &locationAreaRes)
	if err != nil {
		fmt.Println("Errr unmarshaling: %w", err)
		return LocationArea{}, fmt.Errorf("Error unmarshalling: %w", err)
	}
	return locationAreaRes, nil
}

func GetLocationAreaInfo(url string, cache *pokecache.Cache) (LocationAreaInfo, error) {
	//check cache
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error requesting: %w", err)
			return LocationAreaInfo{}, fmt.Errorf("Error requesting: pokeapi.co/api/v2/location-area/: %w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading body: %w", err)
			return LocationAreaInfo{}, fmt.Errorf("Error reading body: %w", err)
		}
		cache.Add(url, body)
	}

	var locationAreaInfoRes LocationAreaInfo
	err := json.Unmarshal(body, &locationAreaInfoRes)
	if err != nil {
		fmt.Println("Errr unmarshaling: %w", err)
		return LocationAreaInfo{}, fmt.Errorf("Error unmarshalling: %w", err)
	}
	return locationAreaInfoRes, nil
}
