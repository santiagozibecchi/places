package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/places/models"
	"github.com/places/services"
)

func CreatePlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPlace models.Place

	err := json.NewDecoder(r.Body).Decode(&newPlace)
	fmt.Println(newPlace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPlaceName, err := services.CreatePlace(newPlace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al crear el Lugar"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("Lugar '%s' creado con exito!", newPlaceName)
	w.Write([]byte(response))
}

func GetPlaces(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	places, err := services.GetAllPlaces()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al traer todos los lugares"))
		return
	}

	w.WriteHeader(http.StatusFound)
	/*
	se encarga de convertir el slice de películas (movies) a JSON
	y escribirlo en el http.ResponseWriter para enviar la 
	respuesta al cliente. 
	*/
	json.NewEncoder(w).Encode(places)

}

func DeletePlace(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
    id := vars["id"]

	deletedPlaceName, err := services.DeleteByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al eliminar el Lugar"))
		return
	}

    w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Lugar '%s' eliminado con exito!", deletedPlaceName)
	fmt.Println(response)
	w.Write([]byte(response))
}