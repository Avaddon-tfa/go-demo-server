package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var Movies []Movie

func main() {
	r := mux.NewRouter()

	Movies = append(Movies, Movie{ID: "1", Isbn: "123456", Title: "Movie 1", Director: &Director{FirstName: "John", LastName: "Doe"}})
	Movies = append(Movies, Movie{ID: "2", Isbn: "789012", Title: "Movie 2", Director: &Director{FirstName: "Jane", LastName: "Smith"}})
	Movies = append(Movies, Movie{ID: "3", Isbn: "783042", Title: "Movie 3", Director: &Director{FirstName: "Michael", LastName: "Black"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	addr := ":8000"
	fmt.Printf("Starting Server at :%s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
