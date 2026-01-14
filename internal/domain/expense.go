package domain

import "time"

type Expense struct {
	ID        int64    `json:"id"`
	Title     string    `json:"title"`
	Amount    int64     `json:"amount"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
