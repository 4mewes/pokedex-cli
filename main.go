package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/4mewes/pokedex/internal/pokecache"
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

func getPokemonInfo(url string, cache *pokecache.Cache) (pokemonInfo, error) {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("Error requesting: %w", err)
			return pokemonInfo{}, fmt.Errorf("Error requesting: pokeapi.co/api/v2/pokemon/: %w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading body: %w", err)
			return pokemonInfo{}, fmt.Errorf("Error reading body: %w", err)
		}
		cache.Add(url, body)
	}

	var pokemonInfoRes pokemonInfo
	err := json.Unmarshal(body, &pokemonInfoRes)
	if err != nil {
		fmt.Println("Errr unmarshaling: %w", err)
		return pokemonInfo{}, fmt.Errorf("Error unmarshalling: %w", err)
	}
	return pokemonInfoRes, nil
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

func commandCatch(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("missing required parameter: pokemon name")
		return nil
	}
	pokemon := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon + "/"

	pokemonInfoRes, err := getPokemonInfo(url, conf.cache)
	if err != nil {
		fmt.Println("error in getPokemonInfo: %w", err)
		return fmt.Errorf("error in getPokemonInfo: %w", err)
	}

	baseExperience := pokemonInfoRes.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if rand.Intn(MaxBaseExp) >= baseExperience {
		if conf.pokedex == nil {
			conf.pokedex = make(map[string]pokemonInfo)
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
	pokedex  map[string]pokemonInfo
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

type pokemonInfo struct {
	Abilities              []Abilities     `json:"abilities,omitempty"`
	BaseExperience         int             `json:"base_experience,omitempty"`
	Cries                  Cries           `json:"cries,omitempty"`
	Forms                  []Forms         `json:"forms,omitempty"`
	GameIndices            []GameIndices   `json:"game_indices,omitempty"`
	Height                 int             `json:"height,omitempty"`
	HeldItems              []HeldItems     `json:"held_items,omitempty"`
	Id                     int             `json:"id,omitempty"`
	IsDefault              bool            `json:"is_default,omitempty"`
	LocationAreaEncounters string          `json:"location_area_encounters,omitempty"`
	Moves                  []Moves         `json:"moves,omitempty"`
	Name                   string          `json:"name,omitempty"`
	Order                  int             `json:"order,omitempty"`
	PastAbilities          []PastAbilities `json:"past_abilities,omitempty"`
	PastTypes              []interface{}   `json:"past_types,omitempty"`
	Species                Species         `json:"species,omitempty"`
	Sprites                Sprites         `json:"sprites,omitempty"`
	Stats                  []Stats         `json:"stats,omitempty"`
	Types                  []Types         `json:"types,omitempty"`
	Weight                 int             `json:"weight,omitempty"`
}

type Ability struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Abilities struct {
	Ability  Ability `json:"ability,omitempty"`
	IsHidden bool    `json:"is_hidden,omitempty"`
	Slot     int     `json:"slot,omitempty"`
}

type Cries struct {
	Latest string `json:"latest,omitempty"`
	Legacy string `json:"legacy,omitempty"`
}

type Forms struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type GameIndices struct {
	GameIndex int     `json:"game_index,omitempty"`
	Version   Version `json:"version,omitempty"`
}

type Item struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type HeldItems struct {
	Item           Item             `json:"item,omitempty"`
	VersionDetails []VersionDetails `json:"version_details,omitempty"`
}

type Move struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type MoveLearnMethod struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type VersionGroup struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at,omitempty"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method,omitempty"`
	Order           interface{}     `json:"order,omitempty"`
	VersionGroup    VersionGroup    `json:"version_group,omitempty"`
}

type Moves struct {
	Move                Move                  `json:"move,omitempty"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details,omitempty"`
}

type Generation struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type PastAbilities struct {
	Abilities  []Abilities `json:"abilities,omitempty"`
	Generation Generation  `json:"generation,omitempty"`
}

type Species struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type DreamWorld struct {
	FrontDefault string      `json:"front_default,omitempty"`
	FrontFemale  interface{} `json:"front_female,omitempty"`
}

type Home struct {
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default,omitempty"`
	FrontShiny   string `json:"front_shiny,omitempty"`
}

type Showdown struct {
	BackDefault      string      `json:"back_default,omitempty"`
	BackFemale       string      `json:"back_female,omitempty"`
	BackShiny        string      `json:"back_shiny,omitempty"`
	BackShinyFemale  interface{} `json:"back_shiny_female,omitempty"`
	FrontDefault     string      `json:"front_default,omitempty"`
	FrontFemale      string      `json:"front_female,omitempty"`
	FrontShiny       string      `json:"front_shiny,omitempty"`
	FrontShinyFemale string      `json:"front_shiny_female,omitempty"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world,omitempty"`
	Home            Home            `json:"home,omitempty"`
	OfficialArtwork OfficialArtwork `json:"official-artwork,omitempty"`
	Showdown        Showdown        `json:"showdown,omitempty"`
}

type RedBlue struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackGray         string `json:"back_gray,omitempty"`
	BackTransparent  string `json:"back_transparent,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontGray        string `json:"front_gray,omitempty"`
	FrontTransparent string `json:"front_transparent,omitempty"`
}

type Yellow struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackGray         string `json:"back_gray,omitempty"`
	BackTransparent  string `json:"back_transparent,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontGray        string `json:"front_gray,omitempty"`
	FrontTransparent string `json:"front_transparent,omitempty"`
}

type GenerationI struct {
	RedBlue RedBlue `json:"red-blue,omitempty"`
	Yellow  Yellow  `json:"yellow,omitempty"`
}

type Crystal struct {
	BackDefault           string `json:"back_default,omitempty"`
	BackShiny             string `json:"back_shiny,omitempty"`
	BackShinyTransparent  string `json:"back_shiny_transparent,omitempty"`
	BackTransparent       string `json:"back_transparent,omitempty"`
	FrontDefault          string `json:"front_default,omitempty"`
	FrontShiny            string `json:"front_shiny,omitempty"`
	FrontShinyTransparent string `json:"front_shiny_transparent,omitempty"`
	FrontTransparent      string `json:"front_transparent,omitempty"`
}

type Gold struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackShiny        string `json:"back_shiny,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontTransparent string `json:"front_transparent,omitempty"`
}

type Silver struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackShiny        string `json:"back_shiny,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontTransparent string `json:"front_transparent,omitempty"`
}

type GenerationIi struct {
	Crystal Crystal `json:"crystal,omitempty"`
	Gold    Gold    `json:"gold,omitempty"`
	Silver  Silver  `json:"silver,omitempty"`
}

type Emerald struct {
	FrontDefault string `json:"front_default,omitempty"`
	FrontShiny   string `json:"front_shiny,omitempty"`
}

type FireredLeafgreen struct {
	BackDefault  string `json:"back_default,omitempty"`
	BackShiny    string `json:"back_shiny,omitempty"`
	FrontDefault string `json:"front_default,omitempty"`
	FrontShiny   string `json:"front_shiny,omitempty"`
}

type RubySapphire struct {
	BackDefault  string `json:"back_default,omitempty"`
	BackShiny    string `json:"back_shiny,omitempty"`
	FrontDefault string `json:"front_default,omitempty"`
	FrontShiny   string `json:"front_shiny,omitempty"`
}

type GenerationIii struct {
	Emerald          Emerald          `json:"emerald,omitempty"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen,omitempty"`
	RubySapphire     RubySapphire     `json:"ruby-sapphire,omitempty"`
}

type DiamondPearl struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackFemale       string `json:"back_female,omitempty"`
	BackShiny        string `json:"back_shiny,omitempty"`
	BackShinyFemale  string `json:"back_shiny_female,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type HeartgoldSoulsilver struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackFemale       string `json:"back_female,omitempty"`
	BackShiny        string `json:"back_shiny,omitempty"`
	BackShinyFemale  string `json:"back_shiny_female,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type Platinum struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackFemale       string `json:"back_female,omitempty"`
	BackShiny        string `json:"back_shiny,omitempty"`
	BackShinyFemale  string `json:"back_shiny_female,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type GenerationIv struct {
	DiamondPearl        DiamondPearl        `json:"diamond-pearl,omitempty"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver,omitempty"`
	Platinum            Platinum            `json:"platinum,omitempty"`
}

type ScarletViolet struct {
	FrontDefault string      `json:"front_default,omitempty"`
	FrontFemale  interface{} `json:"front_female,omitempty"`
}

type GenerationIx struct {
	ScarletViolet ScarletViolet `json:"scarlet-violet,omitempty"`
}

type Animated struct {
	BackDefault      string `json:"back_default,omitempty"`
	BackFemale       string `json:"back_female,omitempty"`
	BackShiny        string `json:"back_shiny,omitempty"`
	BackShinyFemale  string `json:"back_shiny_female,omitempty"`
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type BlackWhite struct {
	Animated         Animated `json:"animated,omitempty"`
	BackDefault      string   `json:"back_default,omitempty"`
	BackFemale       string   `json:"back_female,omitempty"`
	BackShiny        string   `json:"back_shiny,omitempty"`
	BackShinyFemale  string   `json:"back_shiny_female,omitempty"`
	FrontDefault     string   `json:"front_default,omitempty"`
	FrontFemale      string   `json:"front_female,omitempty"`
	FrontShiny       string   `json:"front_shiny,omitempty"`
	FrontShinyFemale string   `json:"front_shiny_female,omitempty"`
}

type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white,omitempty"`
}

type OmegarubyAlphasapphire struct {
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type XY struct {
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire,omitempty"`
	XY                     XY                     `json:"x-y,omitempty"`
}

type Icons struct {
	FrontDefault string      `json:"front_default,omitempty"`
	FrontFemale  interface{} `json:"front_female,omitempty"`
}

type UltraSunUltraMoon struct {
	FrontDefault     string `json:"front_default,omitempty"`
	FrontFemale      string `json:"front_female,omitempty"`
	FrontShiny       string `json:"front_shiny,omitempty"`
	FrontShinyFemale string `json:"front_shiny_female,omitempty"`
}

type GenerationVii struct {
	Icons             Icons             `json:"icons,omitempty"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon,omitempty"`
}

type BrilliantDiamondShiningPearl struct {
	FrontDefault string      `json:"front_default,omitempty"`
	FrontFemale  interface{} `json:"front_female,omitempty"`
}

type GenerationViii struct {
	BrilliantDiamondShiningPearl BrilliantDiamondShiningPearl `json:"brilliant-diamond-shining-pearl,omitempty"`
	Icons                        Icons                        `json:"icons,omitempty"`
}

type Versions struct {
	GenerationI    GenerationI    `json:"generation-i,omitempty"`
	GenerationIi   GenerationIi   `json:"generation-ii,omitempty"`
	GenerationIii  GenerationIii  `json:"generation-iii,omitempty"`
	GenerationIv   GenerationIv   `json:"generation-iv,omitempty"`
	GenerationIx   GenerationIx   `json:"generation-ix,omitempty"`
	GenerationV    GenerationV    `json:"generation-v,omitempty"`
	GenerationVi   GenerationVi   `json:"generation-vi,omitempty"`
	GenerationVii  GenerationVii  `json:"generation-vii,omitempty"`
	GenerationViii GenerationViii `json:"generation-viii,omitempty"`
}

type Sprites struct {
	BackDefault      string   `json:"back_default,omitempty"`
	BackFemale       string   `json:"back_female,omitempty"`
	BackShiny        string   `json:"back_shiny,omitempty"`
	BackShinyFemale  string   `json:"back_shiny_female,omitempty"`
	FrontDefault     string   `json:"front_default,omitempty"`
	FrontFemale      string   `json:"front_female,omitempty"`
	FrontShiny       string   `json:"front_shiny,omitempty"`
	FrontShinyFemale string   `json:"front_shiny_female,omitempty"`
	Other            Other    `json:"other,omitempty"`
	Versions         Versions `json:"versions,omitempty"`
}

type Stat struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Stats struct {
	BaseStat int  `json:"base_stat,omitempty"`
	Effort   int  `json:"effort,omitempty"`
	Stat     Stat `json:"stat,omitempty"`
}

type Type struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Types struct {
	Slot int  `json:"slot,omitempty"`
	Type Type `json:"type,omitempty"`
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
