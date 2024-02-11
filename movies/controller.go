package movies

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/abdulbari149/gomovies/utils"
	"github.com/gorilla/mux"
)

type MovieController struct {
	movieRepo MovieRepo
}

func (mc *MovieController) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := mc.movieRepo.ListMovies()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func (mc *MovieController) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	movie, err := mc.movieRepo.GetMovie(id)

	if err != nil {
		utils.SendError(w, http.StatusNotFound, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (mc *MovieController) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, errors.New("Invalid request payload"))
		return
	}

	movie, err := mc.movieRepo.CreateMovie(data)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func (mc *MovieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, errors.New("Invalid request payload"))
		return
	}

	updatedMovie, err := mc.movieRepo.UpdateMovie(id, data)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMovie)
}

func (mc *MovieController) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := mc.movieRepo.DeleteMovie(id)

	if err != nil {
		utils.SendError(w, http.StatusNotFound, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
