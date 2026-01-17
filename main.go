package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/4mewes/pokedex/internal/pokecache"
	"github.com/4mewes/pokedex/internal/pokeapi"
)

const MaxBaseExp = 255

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

func commandMap(conf *config, args ...string) error {
	var url string
	if conf.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20" //pokeapi default
	} else {
		url = conf.next
	}
	locationAreaRes, err := pokeapi.GetLocationArea(url, conf.cache)
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

	locationAreaRes, err := pokeapi.GetLocationArea(url, conf.cache)
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
	locationAreaInfoRes, err := pokeapi.GetLocationAreaInfo(url, conf.cache)

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

func commandCatch(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("missing required parameter: pokemon name")
		return nil
	}
	pokemon := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon + "/"

	pokemonInfoRes, err := pokeapi.GetPokemonInfo(url, conf.cache)
	if err != nil {
		fmt.Println("error in GetPokemonInfo: %w", err)
		return fmt.Errorf("error in GetPokemonInfo: %w", err)
	}

	baseExperience := pokemonInfoRes.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if rand.Intn(MaxBaseExp) >= baseExperience {
		if conf.pokedex == nil {
			conf.pokedex = make(map[string]pokeapi.PokemonInfo)
		}
		conf.pokedex[pokemon] = pokemonInfoRes
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}

func printPokemonInfoFromPokedex(conf *config, pokemonName string) error {
	pokemon, ok := conf.pokedex[pokemonName]
	if !ok {
		fmt.Printf("%s is not in your pokedex!\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for i := range len(pokemon.Stats) {
		fmt.Printf("  - %s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}

	fmt.Printf("Types: \n")
	for i := range len(pokemon.Types) {
		fmt.Printf("  - %s\n", pokemon.Types[i].Type.Name)
	}
	return nil
}

func commandInspect(conf *config, args ...string) error {
	if conf.pokedex == nil || len(conf.pokedex) == 0 {
		fmt.Println("Your pokedex is empty. Catch some pokemon first!")
		return nil
	}
	if len(args) == 0 {
		fmt.Println("please provide a pokemon name to inspect")
		return nil
	}
	pokemonName := args[0]
	return printPokemonInfoFromPokedex(conf, pokemonName)
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
	pokedex  map[string]pokeapi.PokemonInfo
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
		"catch": {
			name:        "catch",
			description: "attampt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect your pokedex",
			callback:    commandInspect,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	conf := config{}
	conf.cache = pokecache.NewCache(5 * time.Second)
	conf.pokedex = make(map[string]pokeapi.PokemonInfo)

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
