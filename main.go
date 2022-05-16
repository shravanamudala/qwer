package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movie

func main() {
	r := mux.NewRouter()
	movies = append(movies, movie{ID: "1", Isbn: "3333", Title: "RRR", Director: &Director{Firstname: "Rajamouli", Lastname: "ss"}})
	movies = append(movies, movie{ID: "2", Isbn: "1234", Title: "KGF", Director: &Director{Firstname: "Prashant", Lastname: "neel"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("starting server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mm movie
	_ = json.NewDecoder(r.Body).Decode(&mm)
	mm.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, mm)
	json.NewEncoder(w).Encode(movies)

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var mm movie
			_ = json.NewDecoder(r.Body).Decode(&mm)
			mm.ID = params["id"]
			movies = append(movies, mm)

		}
	}
	json.NewEncoder(w).Encode(movies)
}
