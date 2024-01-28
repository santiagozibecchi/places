package models

type RequestPlaceLocation struct {
	Description    string `json:"description"`
	Address    	   string `json:"address"`
	EndTime        string `json:"end_time"`
	Kind           string `json:"kind"`
	PlaceName      string `json:"place_name"`
	StartTime      string `json:"start_time"`
	Country        string `json:"country"`
	Location       string `json:"location"`
}

type Place struct {
	PlaceId           int `json:"place_id"`
	LocationId        int `json:"location_id"`
	Description    string `json:"description"`
	EndTime        string `json:"end_time"`
	Kind           string `json:"kind"`
	LatestViews       int `json:"latest_views "`
	PlaceName      string `json:"place_name"`
	StartTime      string `json:"start_time"`
	TotalView         int `json:"total_view"`
	Address    	   string `json:"address"`
}
