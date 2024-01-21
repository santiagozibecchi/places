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

// TODO: utils.go
func StringToInt(str string) int {

	resutl, err := strconv.Atoi(str)
	if err != nil {
		// TODO!
		// nunca estoy atrapando el panic!
		fmt.Printf("Y esto que eesss°?! %v", str)
		panic(err)
	}
	return resutl
}


func CreateCommentInPlaceByUser(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
    placeId, userId := StringToInt(vars["placeId"]), StringToInt(vars["userId"])

	var newComment models.Comment
	err := json.NewDecoder(r.Body).Decode(&newComment)
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