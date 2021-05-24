package apistructures

type Item struct {
	RawID        int          `json:"id"`
	FileName     string       `json:"file-name"`
	Names        Name         `json:"name"`
	Availability Availability `json:"availability"`
	Shadow       string       `json:"shadow,omitempty"`
	Speed        string       `json:"speed,omitempty"`
	Price        int          `json:"price"`
	PriceCj      int          `json:"price-cj,omitempty"`
	PriceFlick   int          `json:"price-flick,omitempty"`
	CatchPhrase  string       `json:"catch-phrase,omitempty"`
	MuseumPhrase string       `json:"museum-phrase,omitempty"`
	ImageURI     string       `json:"image_uri"`
	IconURI      string       `json:"icon_uri"`
	Type         string       `json:"item_type"`
}

// SUB STRUCTS

// Name struct handles mapping for languages
type Name struct {
	UsEn string `json:"name-USen"`
	EuEn string `json:"name-EUen"`
	EuDe string `json:"name-EUde"`
	EuEs string `json:"name-EUes"`
	UsEs string `json:"name-USes"`
	EuFr string `json:"name-EUfr"`
	UsFr string `json:"name-USfr"`
	EuIt string `json:"name-EUit"`
	EuNl string `json:"name-EUnl"`
	NnZh string `json:"name-CNzh"`
	TwZh string `json:"name-TWzh"`
	JpJa string `json:"name-JPja"`
	KrKo string `json:"name-KRko"`
	EuRu string `json:"name-EUru"`
}

// Availability struct
type Availability struct {
	MonthNorthern      string `json:"month-northern"`
	MonthSouthern      string `json:"month-southern"`
	MonthArrayNorthern []int  `json:"month-array-northern"`
	MonthArraySouthern []int  `json:"month-array-southern"`
	Time               string `json:"time"`
	IsAllDay           bool   `json:"isAllDay"`
	IsAllYear          bool   `json:"isAllYear"`
	Location           string `json:"location"`
	Rarity             string `json:"rarity"`
	TimeArray          []int  `json:"time-array"`
}
