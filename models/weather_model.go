package models

type Weather struct {
	WeatherId               int `json:"weather_id"`
	LocationId              int `json:"location_id"`
	Description         *string `json:"description"`
	Temperature        *float64 `json:"temperature"`
	TemperatureMax     *float64 `json:"temperature_max"`
	TemperatureMin     *float64 `json:"temperature_min"`
}