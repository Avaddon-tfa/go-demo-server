package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Movies {
		if item.ID == params["ID"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range Movies {
		if item.ID == params["ID"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000))
	Movies = append(Movies, movie)
	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		return
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)

	for index, item := range Movies {
		if item.ID == params["ID"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["ID"]
			Movies = append(Movies, movie)
			err := json.NewEncoder(w).Encode(movie)
			if err != nil {
				return
			}
		}
	}
}
