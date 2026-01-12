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

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	fmt.Println("Displays a help message")
	for _, command := range commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func getLocationArea(url string, cache pokecache.Cache) (locationArea, error) {
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

func commandMap(conf *config) error {
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

func commandMapb(conf *config) error {
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

func commandExplore(conf *config) error {
	// TODO: IMPLEMENT
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     string
	previous string
	cache    pokecache.Cache
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
			err := commandMap.callback(&conf)
			if err != nil {
				break
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
