package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/places/services"
)

func GetWeatherByPlaceName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	
    placeName := (vars["placeName"])

	weather, err := services.GetWeatherByPlaceName(placeName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := fmt.Sprintf("Error al buscar el clima del lugar: %v", placeName)
		w.Write([]byte(resp))
		return
	}

	json.NewEncoder(w).Encode(weather)
}

func GetWeatherByLocationId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	
    locationId, err := StringToInt(vars["locationId"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	weather, err := services.GetWeatherByLocationId(locationId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := fmt.Sprintf("Error al buscar el clima con el location_id: %v", locationId)
		w.Write([]byte(resp))
		return
	}

	json.NewEncoder(w).Encode(weather)
}