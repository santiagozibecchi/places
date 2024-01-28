package models

type Location struct {
	LocationId       int `json:"location_id"`
	Country       string `json:"country"`
	Latitude     float64 `json:"latitude"`
	Location     float64 `json:"location"`
	Longitude    float64 `json:"longitude"`
}

type GeocodingPlaceByLocation struct {
	Lng float64
	Lat float64
}
