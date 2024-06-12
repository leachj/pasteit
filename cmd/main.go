package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {

	store := make(map[string]string)

	router := http.NewServeMux()

	router.HandleFunc("POST /paste/", func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		id := uuid.New()
		store[id.String()] = string(b)

		w.Header().Add("Location", fmt.Sprintf("http://localhost:8080/paste/%s", id.String()))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("hello"))
	})

	router.HandleFunc("GET /paste/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		body := store[id]

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe() // Run the http server

}
