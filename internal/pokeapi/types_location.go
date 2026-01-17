package pokeapi

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaInfo struct {
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
