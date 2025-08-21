package dto

type Pokemon struct {
	Abilities              []Ability     `json:"abilities"`
	BaseExperience         int           `json:"base_experience"`
	Cries                  Cries         `json:"cries"`
	Forms                  []Form        `json:"forms"`
	GameIndices            []Index       `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []HeldItem    `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Mfe         `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int           `json:"order"`
	PastAbilities          []PastAbility `json:"past_abilities"`
	PastTypes              []interface{} `json:"past_types"`
	Species                Species       `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Stats                  []Stat        `json:"stats"`
	Types                  []Type        `json:"types"`
	Weight                 int           `json:"weight"`
}

type Ability struct {
	Ability  Ability2 `json:"ability"`
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
}

type Ability2 struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Index struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type HeldItem struct {
	Item           Item            `json:"item"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetail struct {
	Rarity  int      `json:"rarity"`
	Version Version2 `json:"version"`
}

type Version2 struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Mfe struct {
	Move                Move                 `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	Order           interface{}     `json:"order"`
	VersionGroup    VersionGroup    `json:"version_group"`
}

type MoveLearnMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PastAbility struct {
	Abilities  []Ability3 `json:"abilities"`
	Generation Generation `json:"generation"`
}

type Ability3 struct {
	Ability  interface{} `json:"ability"`
	IsHidden bool        `json:"is_hidden"`
	Slot     int         `json:"slot"`
}

type Generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Sprites struct {
	BackDefault      string   `json:"back_default"`
	BackFemale       *string  `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  *string  `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      *string  `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale *string  `json:"front_shiny_female"`
	Other            Other    `json:"other"`
	Versions         Versions `json:"versions"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}

type DreamWorld struct {
	FrontDefault string  `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type Home struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Showdown struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// versões (exemplo simplificado, mas segue o mesmo padrão pra cada geração)
type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationIi   GenerationIi   `json:"generation-ii"`
	GenerationIii  GenerationIii  `json:"generation-iii"`
	GenerationIv   GenerationIv   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVi   GenerationVi   `json:"generation-vi"`
	GenerationVii  GenerationVii  `json:"generation-vii"`
	GenerationViii GenerationViii `json:"generation-viii"`
}

type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  Yellow  `json:"yellow"`
}

type RedBlue struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}

type Yellow struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}

// --- Generation II ---
type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Silver  `json:"silver"`
}

type Crystal struct {
	BackDefault           string `json:"back_default"`
	BackShiny             string `json:"back_shiny"`
	BackShinyTransparent  string `json:"back_shiny_transparent"`
	BackTransparent       string `json:"back_transparent"`
	FrontDefault          string `json:"front_default"`
	FrontShiny            string `json:"front_shiny"`
	FrontShinyTransparent string `json:"front_shiny_transparent"`
	FrontTransparent      string `json:"front_transparent"`
}

type Gold struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}

type Silver struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}

// --- Generation III ---
type GenerationIii struct {
	Emerald          Emerald          `json:"emerald"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
	RubySapphire     RubySapphire     `json:"ruby-sapphire"`
}

type Emerald struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type FireredLeafgreen struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type RubySapphire struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

// --- Generation IV ---
type GenerationIv struct {
	DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
	Platinum            Platinum            `json:"platinum"`
}

type DiamondPearl struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type HeartgoldSoulsilver struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type Platinum struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// --- Generation V ---
type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}

type BlackWhite struct {
	Animated         Animated `json:"animated"`
	BackDefault      string   `json:"back_default"`
	BackFemale       *string  `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  *string  `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      *string  `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale *string  `json:"front_shiny_female"`
}

type Animated struct {
	BackDefault      string  `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        string  `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// --- Generation VI ---
type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
	XY                     XY                     `json:"x-y"`
}

type OmegarubyAlphasapphire struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type XY struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// --- Generation VII ---
type GenerationVii struct {
	Icons             Icons             `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}

type Icons struct {
	FrontDefault string  `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type UltraSunUltraMoon struct {
	FrontDefault     string  `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       string  `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// --- Generation VIII ---
type GenerationViii struct {
	Icons Icons2 `json:"icons"`
}

type Icons2 struct {
	FrontDefault string  `json:"front_default"`
	FrontFemale  *string `json:"front_female"`
}

type Stat struct {
	BaseStat int   `json:"base_stat"`
	Effort   int   `json:"effort"`
	Stat     Stat2 `json:"stat"`
}

type Stat2 struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Type struct {
	Slot int   `json:"slot"`
	Type Type2 `json:"type"`
}

type Type2 struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
