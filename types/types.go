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

type Properties struct {
	MapboxID string `json:"mapbox_id"`
	Wikidata string `json:"wikidata"`
	// Otros campos relacionados con las propiedades
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
