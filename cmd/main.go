package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/XoDeR/simple-rest-api-go/pkg/db"
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
	DB := db.Init()
	h := myHandlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("BookSeller API")
	})

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	log.Println("API server is running!")
	http.ListenAndServe(":4000", corsHandler(router))

	return nil
}
