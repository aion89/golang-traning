package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Article asd
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

//Customer
type Customer struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	State string `json:"state"`
}

//Articles asd
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:       "Test Title",
			Description: "Test Description",
			Content:     "Test Content",
		},
		Article{
			Title:       "Test Title 2",
			Description: "Test Description 2",
			Content:     "Test Content 2",
		},
	}

	fmt.Println("Endpoint hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: GET Home Page Endpoint")
	fmt.Fprintf(w, "Endpoint hit: GET Home Page Endpoint")
	fmt.Fprintf(w, "Супер мега приложуха")
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: POST All Articles Endpoint")
	fmt.Fprintf(w, "Endpoint hit: POST All Articles Endpoint")
}

func vovic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Get Vovic")
	vovic := Customer{
		Name:  "Vovic",
		ID:    "1",
		State: "Mudak",
	}
	json.NewEncoder(w).Encode(vovic)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/vovic", vovic).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	err := http.ListenAndServe(":8081", myRouter)
	if err != nil {
		log.Fatalf("Server failed to start with error %v", err)
	}

}

func main() {
	handleRequests()
}
