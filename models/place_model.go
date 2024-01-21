package models

type Place struct {
	PlaceID     int    `json:"place_id"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
	TotalView   string `json:"total_view"`
}
