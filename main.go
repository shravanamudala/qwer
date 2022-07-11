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

type article struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Content   string    `json:"content"`
	CreationTime string   `json:"creationtime"`
}


var articles []article

func main() {
	r := mux.NewRouter()
	articles = append(articles, article{ID: "1", Title: "country", Subtitle: "mera bharat mahaan",Content:"India is my country. new delhi is the capital of india", CreationTime: "11-7-22"})
    articles = append(articles, article{ID: "2", Title: "state", Subtitle: "telangana",Content:"it is a state in india. hyderabad is the capital of telangana",CreationTime: "11-7-22"})
	r.HandleFunc("/articles", getArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	r.HandleFunc("/articles", createArticle).Methods("POST")
	
	fmt.Printf("starting server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}


func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mm article
	_ = json.NewDecoder(r.Body).Decode(&mm)
	mm.ID = strconv.Itoa(rand.Intn(1000000))
	articles = append(articles, mm)
	json.NewEncoder(w).Encode(articles)

}
