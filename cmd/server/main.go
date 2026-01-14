package main

import (
	"log"
	"net/http"
	"personal_expense_tracker/internal/db"
	"personal_expense_tracker/internal/repository"
	"personal_expense_tracker/internal/service"
	handler "personal_expense_tracker/internal/transport/http"
	"time"

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
	defer db.Close()

	expenseRepo := repository.NewExpenseRepostory(db)
	expenseService := service.NewExpenseService(expenseRepo)
	expenseHandler := handler.NewExpenseHandler(expenseService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /expenses", expenseHandler.CreateExpense)
	mux.HandleFunc("GET /expenses", expenseHandler.GetExpenses)
	mux.HandleFunc("GET /expenses/{id}", expenseHandler.GetExpenseById)
	mux.HandleFunc("PATCH /expenses/{id}", expenseHandler.UpdateExpenseById)
	mux.HandleFunc("DELETE /expenses/{id}", expenseHandler.DeleteExpenseById)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
