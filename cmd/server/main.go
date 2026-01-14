package main

import (
	"log"
	"net/http"
	"personal_expense_tracker/internal/db"
	"personal_expense_tracker/internal/repository"

	"github.com/joho/godotenv"
)

func main() {
	// load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("main_error: error loading env vars")
	}

	// connect to DB
	db, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal("main_error: db connection failed with error" + err.Error())
	}

	// create repo instance
	repository.NewExpenseRepostory(db)

	// create service instance

	// create handler instance

	// serve app
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("main_error: failed to listen on port:8080")
	}
	log.Print("listening on port:8080")
}
