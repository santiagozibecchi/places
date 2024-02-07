package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/places/routes"
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
	routes.InitRoutes(router)

	// Cors
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := corsOptions.Handler(router)

	port := ":8080"
	if  err := StartServer(port, handler); err != nil {
		log.Fatalf("Error al inicial el servidor: %v", err)
	}

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