package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db_assets "github.com/mhafidk/warga-sekolah/db"
	"github.com/pressly/goose/v3"
	_ "turso.tech/database/tursogo"
)

func main() {
	db, err := sql.Open("turso", "db/app.db")
	if err != nil {
		log.Fatal("Failed to open the database: ", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	defer db.Close()
	log.Println("Database was opened successfully!")

	goose.SetBaseFS(db_assets.MigrationFS)

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal("Failed to set the goose dialect: ", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	log.Println("Database migrations applied successfully!")

	http.HandleFunc("/ping", ping)

	log.Println("Start the server at port 3000")
	http.ListenAndServe(":3000", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!\n")
}
