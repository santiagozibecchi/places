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

func GetSpecificPlace(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
    placeId := vars["id"]

	place, err := services.GetPlaceById(placeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := fmt.Sprintf("Error al traer el lugar con el id: %v", placeId)
		w.Write([]byte(resp))
		return
	}

	json.NewEncoder(w).Encode(place)
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

func UpdatePlace(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID y los datos actualizados del cuerpo de la solicitud
	vars := mux.Vars(r)
	id := vars["id"]

	var editedPlace models.Place
	err := json.NewDecoder(r.Body).Decode(&editedPlace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedRow, err := services.UpdateByID(id, editedPlace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedRow)
}
