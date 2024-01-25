package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/places/controllers"
)

type setHandlerFunc func(path string, f http.HandlerFunc)

func BuildSetHandleFunc(router *mux.Router, methods ...string) setHandlerFunc {
	return func(path string, f http.HandlerFunc) {
		router.HandleFunc(path, f).Methods(methods...)
	}
}

func InitRoutes(router *mux.Router) {

	Post := BuildSetHandleFunc(router, "POST")
	Get := BuildSetHandleFunc(router, "GET")
	Put := BuildSetHandleFunc(router, "PUT")
	Delete := BuildSetHandleFunc(router, "DELETE")
	
	// Places
	// TODO: id => placeId
	Get("/api/v1/places", controllers.GetPlaces)
	Get("/api/v1/places/{id}", controllers.GetSpecificPlace)
	Post("/api/v1/places", controllers.CreatePlace)
	Delete("/api/v1/places/{id}", controllers.DeletePlace)
	Put("/api/v1/places/{id}", controllers.UpdatePlace)
	// TODO: nice to implement
	// r.HandleFunc("/api/v1/places/{groupBy:(kind}", controllers.HandleFunc)

	// Search Places
	Get("/api/v1/places/placeName/{placeName}", controllers.SearchPlaces)

	// Users
	Get("/api/v1/users", controllers.GetUsers)
	Get("/api/v1/users", controllers.GetUsers)
	Post("/api/v1/users", controllers.CreateUser)
	Delete("/api/v1/users/{userId}", controllers.DeleteUser)
	Put("/api/v1/users/{userId}", controllers.UpdateUser)
	
	// Comments
	Post("/api/v1/place/{placeId}/user/{userId}", controllers.CreateCommentInPlaceByUser)
}