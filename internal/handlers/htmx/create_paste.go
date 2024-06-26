package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func CreatePaste(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Could not parse form.", http.StatusBadRequest)
			return
		}

		data := r.PostForm.Get("text-area")

		id := uuid.New()
		stmt, err := db.Prepare("insert into pastes(id, data) values (?, ?) ")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(id.String(), string(data))
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Location", fmt.Sprintf("http://localhost:8080/paste/%s", id.String()))
		w.WriteHeader(http.StatusMovedPermanently)
		//w.Write([]byte(fmt.Sprintf("<b>new pasteit</b><p>%s</p>", data)))
	}

}
