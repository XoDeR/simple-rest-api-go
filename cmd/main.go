package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	myHandlers "github.com/XoDeR/simple-rest-api-go/pkg/handlers"
	"github.com/gorilla/handlers"
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

	router.HandleFunc("/books", myHandlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", myHandlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", myHandlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", myHandlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", myHandlers.DeleteBook).Methods(http.MethodDelete)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Println("API server is running!")
	http.ListenAndServe(":4000", corsHandler(router))

	return nil
}
