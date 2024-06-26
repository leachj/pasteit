package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	handlers "github.com/leachj/pasteit/internal/handlers/htmx"
)

func main() {

	db, err := sql.Open("sqlite3", "./pasteit.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//store := make(map[string]string)

	sqlStmt := `
	create table if not exists pastes (id varchar not null primary key, data text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	fs := http.FileServer(http.Dir("./web"))

	router := http.NewServeMux()

	router.HandleFunc("POST /paste/", handlers.CreatePaste(db))
	router.HandleFunc("GET /paste/{id}", handlers.GetPaste(db))
	// options calls in router
	router.HandleFunc("OPTIONS /*", handlers.Options())

	router.Handle("/", fs)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe() // Run the http server

}
