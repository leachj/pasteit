package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func GetPaste(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		stmt, err := db.Prepare("select data from pastes where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		var data string
		err = stmt.QueryRow(id).Scan(&data)
		if err != nil {
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
			Data: data,
		}

		w.WriteHeader(http.StatusOK)

		err2 := template.Execute(w, templateData)

		if err2 != nil {
			log.Print(err2)
		}

	}

}
