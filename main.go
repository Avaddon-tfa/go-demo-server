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

	Movies = append(Movies, Movie{ID: "1", Isbn: "428337", Title: "Lord of the Rings", Director: &Director{FirstName: "Peter", LastName: "Jackson"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
