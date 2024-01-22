package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/places/controllers"
	"github.com/places/services"
	
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

func main() {
	err := services.EstablishDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", handleF).Methods("GET")

	// Places
	// TODO: id => placeId
	router.HandleFunc("/api/v1/places", controllers.GetPlaces).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/places/{id}", controllers.GetSpecificPlace).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/places", controllers.CreatePlace).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/places/{id}", controllers.DeletePlace).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/places/{id}", controllers.UpdatePlace).Methods(http.MethodPut)
	// Search Places
	router.HandleFunc("/api/v1/places/placeName/{placeName}", controllers.SearchPlaces).Methods(http.MethodGet)
	// TODO: nice to implement
	// r.HandleFunc("/api/v1/places/{placeName}/{sort:(?:asc|desc|default)}", controllers.HandleFunc)
	// TODO: nice to implement
	// r.HandleFunc("/api/v1/places/{kind}/{sort:(?:asc|desc|default)}", controllers.HandleFunc)

	// Users
	router.HandleFunc("/api/v1/users", controllers.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users/{userId}", controllers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/users/{userId}", controllers.UpdateUser).Methods(http.MethodPut)
	
	// Comments
	router.HandleFunc("/api/v1/place/{placeId}/user/{userId}", controllers.CreateCommentInPlaceByUser).Methods(http.MethodPost)
	
	// Cors

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5432"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := corsOptions.Handler(router)

	port := ":8080"
	if  err := StartServer(port, handler); err != nil {
		log.Fatalf("Error al inicial el servidor: %v", err)
	}

}

func handleF(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola"))
}


func StartServer(port string, router http.Handler) error {
	server := &http.Server{
		Handler: router,
		Addr: port,
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}

	fmt.Printf("Iniciando servidor en el puerto: %s ..\n", port)

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("error iniciando el servidor %v", err)
	}

	return nil
}