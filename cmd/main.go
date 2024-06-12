package main

import (
	"net/http"

	"github.com/leachj/pasteit/internal/handlers"
)

func main() {

	store := make(map[string]string)

	router := http.NewServeMux()

	router.HandleFunc("POST /paste/", handlers.CreatePaste(&store))
	router.HandleFunc("GET /paste/{id}", handlers.GetPaste((&store)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe() // Run the http server

}
