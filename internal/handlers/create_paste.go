package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func CreatePaste(store *map[string]string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		id := uuid.New()
		(*store)[id.String()] = string(b)

		w.Header().Add("Location", fmt.Sprintf("http://localhost:8080/paste/%s", id.String()))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("hello"))
	}

}
