package handlers

import (
	"log"
	"net/http"
)

func GetPaste(store *map[string]string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		body, ok := (*store)[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("Paste with id %s not found", id)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}

}
