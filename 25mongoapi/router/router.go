package router

import (
	"github.com/gorilla/mux"
	"github.com/nishaa007/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllmovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.Markaswatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.Deleteonemovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.Deleteallmovie).Methods("DELETE")

	return router
}
