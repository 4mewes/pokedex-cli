package pokeapi

type PokemonInfo struct {
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
