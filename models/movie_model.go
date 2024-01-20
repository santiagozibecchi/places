package models

type Movie struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Year int `json:"year"`
	Gender string `json:"gender"`
	Adquired string `json:"adquired"`
	Stock int `json:"stock"`
	Price float64 `json:"price"`
}