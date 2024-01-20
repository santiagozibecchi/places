package models

type Place struct {
	PlaceID     int    `json:"place_id"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Description string `json:"description"`
}
