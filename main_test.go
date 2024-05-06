package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Create a test Movies array
var testMovies = []Movie{
	{ID: "1", Isbn: "123456", Title: "Movie 1", Director: &Director{FirstName: "John", LastName: "Doe"}},
	{ID: "1", Isbn: "123456", Title: "Movie 1", Director: &Director{FirstName: "John", LastName: "Doe"}},
}

func TestGetMovies(t *testing.T) {
	// Backup the original data and restore after this test
	originalMovies := Movies
	defer func() {
		Movies = originalMovies
	}()

	Movies = testMovies

	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(getMovies)

	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v but expected %v", status, http.StatusOK)
	}

	var movies []Movie
	err = json.Unmarshal(response.Body.Bytes(), &movies)

	if err != nil {
		t.Fatal(err)
	}

	if len(movies) != len(testMovies) {
		t.Errorf("Expected %v movies, got %v", len(testMovies), len(movies))
	}
}

// You need to implement other test functions for `deleteMovie`, `getMovie`, `createMovie`, and `updateMovie` in a similar way.
