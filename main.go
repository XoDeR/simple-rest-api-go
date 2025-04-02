package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux" // routing
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("BookSeller API")
	})

	log.Println("API server is running!")
	http.ListenAndServe(":4000", router)
}
