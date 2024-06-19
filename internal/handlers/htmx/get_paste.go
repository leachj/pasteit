package handlers

import (
	"log"
	"net/http"
	"text/template"
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

		// load template from file
		template, err := template.ParseFiles("internal/templates/view.tmpl")
		if err != nil {
			// bugger
		}
		templateData := struct {
			Data string
		}{
			Data: body,
		}

		w.WriteHeader(http.StatusOK)

		err2 := template.Execute(w, templateData)

		if err2 != nil {
			log.Print(err2)
		}

	}

}
