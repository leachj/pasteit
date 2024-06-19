package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func CreatePaste(store *map[string]string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Could not parse form.", http.StatusBadRequest)
			return
		}

		data := r.PostForm.Get("text-area")
		id := uuid.New()
		(*store)[id.String()] = string(data)

		w.Header().Add("Location", fmt.Sprintf("http://localhost:8080/paste/%s", id.String()))
		w.WriteHeader(http.StatusMovedPermanently)
		//w.Write([]byte(fmt.Sprintf("<b>new pasteit</b><p>%s</p>", data)))
	}

}
