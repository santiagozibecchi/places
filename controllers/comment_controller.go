package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/places/models"
	"github.com/places/services"
)

func StringToInt(value string) (int, error) {
    intValue, err := strconv.Atoi(value)
    if err != nil {
        return 0, fmt.Errorf("Error al convertir '%s' a entero: %v", value, err)
    }
    return intValue, nil
}


func CreateCommentInPlaceByUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	
    placeId, err := StringToInt(vars["placeId"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	userId, err := StringToInt(vars["userId"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	var newComment models.Comment

	err = json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := services.CreateComment(placeId, userId, newComment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		resp := fmt.Sprintf("Error al crear el commentario con el placeId: %v y userId: %v", placeId, userId)
		w.Write([]byte(resp))
		return
	}

	json.NewEncoder(w).Encode(comment)
}