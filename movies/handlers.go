package movies

import (
	"github.com/gorilla/mux"
)

func InitHandlers(r *mux.Router) {

	mc := &MovieController{
		movieRepo: &MovieRepoImpl{},
	}

	mr := r.PathPrefix("/movies").Subrouter().StrictSlash(true)

	mr.HandleFunc("/", mc.GetMovies).Methods("GET")

	mr.HandleFunc("/{id}", mc.GetMovie).Methods("GET")

	mr.HandleFunc("/", mc.CreateMovie).Methods("POST")

	mr.HandleFunc("/{id}", mc.UpdateMovie).Methods("PUT")

	mr.HandleFunc("/{id}", mc.DeleteMovie).Methods("DELETE")
}
