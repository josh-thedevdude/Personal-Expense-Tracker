# Expense Tracker
A basic crud app for understanding Go REST.

### Folder Structure

```go
expense-tracker/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   └── expense.go
│   │
│   ├── repository/
│   │   └── expense_repository.go
│   │
│   ├── service/
│   │   └── expense_service.go
│   │
│   ├── transport/
│   │   └── http/
│   │       └── expense_handler.go
│   │
│   └── db/
│       └── postgres.go
│
├── go.mod
└── README.md

```