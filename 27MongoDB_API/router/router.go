package router

import (
	"github.com/ChaharSC03/mongodbapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMyMoviesController).Methods("GET")

	r.HandleFunc("/api/movie", controller.CreateMovieController).Methods("POST")

	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatchedController).Methods("PUT")

	r.HandleFunc("/api/movie/{id}", controller.DeleteMovieController).Methods("DELETE")

	r.HandleFunc("/api/deleteall", controller.DeleteAllMoviesController).Methods("DELETE")
	return r

}
