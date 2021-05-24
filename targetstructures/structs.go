package targetstructures

import "time"

type Output struct {
	Time      time.Time          `json:"time_generated"`
	All       map[string]Item    `json:"all"`
	Bugs      map[string]Item    `json:"bugs"`
	Fish      map[string]Item    `json:"fish"`
	Sea       map[string]Item    `json:"sea"`
	Leaving   ItemTypeHemisphere `json:"leaving"`
	New       ItemTypeHemisphere `json:"new"`
	Available ItemTypeHemisphere `json:"available"`
}

type ItemTypeHemisphere struct {
	All      map[string]Item `json:"all"`
	Northern TypedItems      `json:"northern"`
	Southern TypedItems      `json:"southern"`
}

type TypedItems struct {
	All  map[string]Item `json:"all"`
	Bugs map[string]Item `json:"bugs"`
	Fish map[string]Item `json:"fish"`
	Sea  map[string]Item `json:"sea"`
}

// Item is a target structure
type Item struct {
	ID         string     `json:"id"`
	Attributes Attributes `json:"attributes"`
	Meta       Meta       `json:"meta"`
	Converted  bool       `json:"converted"`
}

type Has struct {
	Price        bool `json:"price"`
	Shadow       bool `json:"shadow"`
	Speed        bool `json:"speed"`
	Rarity       bool `json:"rarity"`
	Location     bool `json:"location"`
	Availability bool `json:"availability"`
}

type IsHemisphere struct {
	New       bool `json:"new"`
	Leaving   bool `json:"leaving"`
	Available bool `json:"available"`
}

type Is struct {
	Northern IsHemisphere `json:"northern"`
	Southern IsHemisphere `json:"southern"`
}

type Meta struct {
	Time time.Time `json:"time_generated"`
	Has  Has       `json:"has"`
	Is   Is        `json:"is"`
}

type Attributes struct {
	Titles       Safe         `json:"titles"`
	URIS         Uris         `json:"uris"`
	Type         TypeMeta     `json:"type"`
	Prices       Prices       `json:"prices"`
	Phrases      Phrases      `json:"phrases"`
	Images       Images       `json:"images"`
	Availability Availability `json:"availability"`
	Shadow       string       `json:"shadow"`
	Speed        string       `json:"speed"`
}

// Titles
type Safe struct {
	Original string `json:"original"`
	Safe     string `json:"safe"`
}

// URI data
type Uris struct {
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

// Type provides
type TypeMeta struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

// Prices provides
type Prices struct {
	Store int `json:"store"`
	Cj    int `json:"cj"`
	Flick int `json:"flick"`
}

// Phrases provides
type Phrases struct {
	Capture Safe `json:"capture"`
	Museum  Safe `json:"museum"`
}

type Image struct {
	Direct string `json:"direct"`
	Local  string `json:"local"`
}

// Images provide
type Images struct {
	Thumb Image `json:"thumb"`
	Main  Image `json:"main"`
}

// Hemisphere provides hemisphere specific data
type Hemisphere struct {
	Always    bool    `json:"always"`
	Text      string  `json:"text"`
	Ranges    string  `json:"ranges"`
	Array     []int   `json:"array"`
	Sequences [][]int `json:"sequences"`
}

// Months provides structured data for Availability
type Months struct {
	Always   bool       `json:"always"`
	Northern Hemisphere `json:"northern"`
	Southern Hemisphere `json:"southern"`
}

// Times provides times based Availability data
type Times struct {
	Always bool   `json:"always"`
	Text   string `json:"text"`
	Array  []int  `json:"array"`
}

// Availability provides
type Availability struct {
	Location string `json:"location"`
	Rarity   string `json:"rarity"`
	Months   Months `json:"months"`
	Times    Times  `json:"times"`
}
