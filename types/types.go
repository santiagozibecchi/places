package types

type PlaceQueryParams struct {
	Sort   string
	Kind   string
	Country string
}

type MapboxGeocodingPlace struct {
	Id string `json:"id"`
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
	Country string `json:"country"`
}

// MAPBOX RESPONSE
type Properties struct {
	MapboxID string `json:"mapbox_id"`
	Wikidata string `json:"wikidata"`
}

type Context struct {
	ID        string `json:"id"`
	MapboxID  string `json:"mapbox_id"`
	Wikidata  string `json:"wikidata"`
	ShortCode string `json:"short_code"`
	TextES    string `json:"text_es"`
	LanguageES string `json:"language_es"`
	Text      string `json:"text"`
	Language  string `json:"language"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Feature struct {
	ID               string     `json:"id"`
	Type             string     `json:"type"`
	PlaceType        []string   `json:"place_type"`
	Relevance        float64    `json:"relevance"`
	Properties       Properties `json:"properties"`
	TextES           string     `json:"text_es"`
	LanguageES       string     `json:"language_es"`
	PlaceNameES      string     `json:"place_name_es"`
	Text             string     `json:"text"`
	Language         string     `json:"language"`
	PlaceName        string     `json:"place_name"`
	Bbox             []float64  `json:"bbox"`
	Center           []float64  `json:"center"`
	Geometry         Geometry   `json:"geometry"`
	Context          []Context  `json:"context"`
}

// WEATHER API RESPONSE
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Rain struct {
	OneHour float64 `json:"1h"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherResponse struct {
	Coord       Coord    `json:"coord"`
	Weather     []Weather `json:"weather"`
	Base        string   `json:"base"`
	Main        Main     `json:"main"`
	Visibility  int      `json:"visibility"`
	Wind        Wind     `json:"wind"`
	Rain        Rain     `json:"rain"`
	Clouds      Clouds   `json:"clouds"`
	DT          int64    `json:"dt"`
	Sys         Sys      `json:"sys"`
	Timezone    int      `json:"timezone"`
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Cod         int      `json:"cod"`
}
