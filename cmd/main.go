package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/XoDeR/simple-rest-api-go/pkg/handlers"
	"github.com/gorilla/mux" // routing
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func run() error {
	router := mux.NewRouter()

	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("BookSeller API")
	})

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)

	log.Println("API server is running!")
	http.ListenAndServe(":4000", router)

	return nil
}
