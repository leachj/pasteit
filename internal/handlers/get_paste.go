package handlers

import (
	"net/http"
)

func GetPaste(store *map[string]string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		body := (*store)[id]

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}

}
