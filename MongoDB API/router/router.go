package router

import (
	"github.com/AlperSeyman/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", controller.GetOneMovie).Methods("GET")
	router.HandleFunc("/movie", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/movies", controller.DeleteAllMovie).Methods("DELETE")
	router.HandleFunc("movie/{id}", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
