package main

import (
	"net/http"

	handlers "github.com/leachj/pasteit/internal/handlers/htmx"
)

func main() {

	store := make(map[string]string)

	fs := http.FileServer(http.Dir("./web"))

	router := http.NewServeMux()

	router.HandleFunc("POST /paste/", handlers.CreatePaste(&store))
	router.HandleFunc("GET /paste/{id}", handlers.GetPaste((&store)))
	// options calls in router
	router.HandleFunc("OPTIONS /*", handlers.Options())

	router.Handle("/", fs)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe() // Run the http server

}
