package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/places/models"
	"github.com/places/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Println(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUserName, userId, err := services.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al crear el nuevo usuario"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("Usuario '%s' creado con exito! Con el ID: %v", newUserName, userId)
	w.Write([]byte(response))
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	users, err := services.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error al traer todos los usuarios"))
		return
	}

	w.WriteHeader(http.StatusFound)
	/*
	se encarga de convertir el slice de películas (movies) a JSON
	y escribirlo en el http.ResponseWriter para enviar la 
	respuesta al cliente. 
	*/
	json.NewEncoder(w).Encode(users)

}

func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
    userId := vars["userId"]

	deletedUserName, err := services.DeleteUserByID(userId)

	if err != nil {
		// Error Interno a la consulta SQL
		http.Error(w, err.Error(), http.StatusInternalServerError)

		// Error personalizado
		resp := fmt.Sprintf("Error al eliminar el usuario con el Id %v", userId)
		w.Write([]byte(resp))
		return
	}

    w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Usuario: '%s' con el id: '%s' eliminado con exito!", deletedUserName, userId)
	w.Write([]byte(response))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	var editedUser models.User

	err := json.NewDecoder(r.Body).Decode(&editedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedRow, err := services.UpdateUserByID(userId, editedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		
		resp := fmt.Sprintf("Error al actualziar el usuario con el Id %v", userId)
		w.Write([]byte(resp))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedRow)
}
