package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/4mewes/pokedex/internal/pokecache"
)

func commandExit(conf *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, args ...string) error {
	fmt.Println("Displays a help message")
	for _, command := range commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func getLocationArea(url string, cache *pokecache.Cache) (locationArea, error) {
	//check cache
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error requesting: %w", err)
			return locationArea{}, fmt.Errorf("Error requesting: pokeapi.co/api/v2/location-area/: %w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading body: %w", err)
			return locationArea{}, fmt.Errorf("Error reading body: %w", err)
		}
		cache.Add(url, body)
	}

	var locationAreaRes locationArea
	err := json.Unmarshal(body, &locationAreaRes)
	if err != nil {
		fmt.Println("Errr unmarshaling: %w", err)
		return locationArea{}, fmt.Errorf("Error unmarshalling: %w", err)
	}
	return locationAreaRes, nil
}

func getLocationAreaInfo(url string, cache *pokecache.Cache) (locationAreaInfo, error) {
	//check cache
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error requesting: %w", err)
			return locationAreaInfo{}, fmt.Errorf("Error requesting: pokeapi.co/api/v2/location-area/: %w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading body: %w", err)
			return locationAreaInfo{}, fmt.Errorf("Error reading body: %w", err)
		}
		cache.Add(url, body)
	}

	var locationAreaInfoRes locationAreaInfo
	err := json.Unmarshal(body, &locationAreaInfoRes)
	if err != nil {
		fmt.Println("Errr unmarshaling: %w", err)
		return locationAreaInfo{}, fmt.Errorf("Error unmarshalling: %w", err)
	}
	return locationAreaInfoRes, nil
}

func commandMap(conf *config, args ...string) error {
	var url string
	if conf.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20" //pokeapi default
	} else {
		url = conf.next
	}
	locationAreaRes, err := getLocationArea(url, conf.cache)
	if err != nil {
		fmt.Println("error in getlocationArea: %w", err)
		return fmt.Errorf("error in getLocationArea: %w", err)
	}

	conf.next = locationAreaRes.Next
	conf.previous = locationAreaRes.Previous
	for _, location := range locationAreaRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(conf *config, args ...string) error {
	var url string
	if conf.previous == "" {
		fmt.Println("you're on the first page.")
		return nil
	} else {
		url = conf.previous
	}

	locationAreaRes, err := getLocationArea(url, conf.cache)
	if err != nil {
		fmt.Println("error in getlocationArea: %w", err)
		return fmt.Errorf("error in getLocationArea: %w", err)
	}

	conf.next = locationAreaRes.Next
	conf.previous = locationAreaRes.Previous
	for _, location := range locationAreaRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(conf *config, args ...string) error {
	var locationAreaName string
	if len(args) == 0 {
		fmt.Println("please provide a location area name arg!")
		return nil
	} else {
		locationAreaName = args[0]
	}

	fmt.Printf("Exploring %s...\n", locationAreaName)
	url := "https://pokeapi.co/api/v2/location-area/" + locationAreaName + "/"
	locationAreaInfoRes, err := getLocationAreaInfo(url, conf.cache)

	if err != nil {
		fmt.Println("error in getlocationAreaInfo: %w", err)
		return fmt.Errorf("error in getLocationAreaInfo: %w", err)
	}
	fmt.Println("Found Pokemon:")
	for _, PokemonEncounters := range locationAreaInfoRes.PokemonEncounters {
		fmt.Printf("- %s\n", PokemonEncounters.Pokemon.Name)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	next     string
	previous string
	cache    *pokecache.Cache
}

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationAreaInfo struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates,omitempty"`
	GameIndex            int                    `json:"game_index,omitempty"`
	Id                   int                    `json:"id,omitempty"`
	Location             Location               `json:"location,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Names                []Names                `json:"names,omitempty"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters,omitempty"`
}

type EncounterMethod struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Version struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type VersionDetails struct {
	Rate    int     `json:"rate,omitempty"`
	Version Version `json:"version,omitempty"`
}

type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method,omitempty"`
	VersionDetails  []VersionDetails `json:"version_details,omitempty"`
}

type Location struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Language struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Names struct {
	Language Language `json:"language,omitempty"`
	Name     string   `json:"name,omitempty"`
}

type Pokemon struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon,omitempty"`
	VersionDetails []VersionDetails `json:"version_details,omitempty"`
}

var commandRegistry = map[string]cliCommand{}

func main() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Welcome to the Pokedex!\nUsage:\n\n",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Explore the Pokemon map, load the next page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore the Pokemon map, load the previous page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List Pokemon in given location",
			callback:    commandExplore,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	conf := config{}
	conf.cache = pokecache.NewCache(5 * time.Second)

	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		userInputRaw := scanner.Text()
		userInput := strings.ToLower(strings.TrimSpace(userInputRaw))
		commands := strings.Fields(userInput)

		commandMap, ok := commandRegistry[commands[0]]

		if ok {
			err := commandMap.callback(&conf, commands[1:]...)
			if err != nil {
				break
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
