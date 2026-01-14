package repository

import (
	"context"
	"database/sql"
	"fmt"
	"personal_expense_tracker/internal/domain"
)

// suggests the blueprint of the repository service.
type ExpenseRepository interface {
	Create(ctx context.Context, e *domain.Expense) error
	GetById(ctx context.Context, id int64) (e domain.Expense, err error)
	Get(ctx context.Context) (e []domain.Expense, err error)
	UpdateById(ctx context.Context, id int64, e *domain.Expense) (expense domain.Expense, err error)
	DeleteById(ctx context.Context, id int64) error
}

// unexported struct to impose logical sepration
type expenseRepository struct {
	db *sql.DB
}

// singleton constructor function
func NewExpenseRepostory(db *sql.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) Create(ctx context.Context, e *domain.Expense) error {
	query := `
		INSERT INTO expenses (title, amount, category)
		VALUES ($1, $2, $3)
		RETURNING id, created_at;
	`
	err := r.db.QueryRowContext(ctx, query, e.Title, e.Amount, e.Category).Scan(&e.ID, &e.CreatedAt)
	if err != nil {
		fmt.Printf("repo_error: %s", err.Error())
		return err
	}

	return nil
}

func (r *expenseRepository) GetById(ctx context.Context, id int64) (e domain.Expense, err error) {
	query := `
		SELECT
			id,
			title,
			amount,
			category,
			created_at
		FROM expenses
		WHERE id = $1;
	`
	err = r.db.QueryRowContext(ctx, query, id).Scan(&e.ID, &e.Title, &e.Amount, &e.Category, &e.CreatedAt)
	if err != nil {
		fmt.Printf("repo_error: %s", err.Error())
		return domain.Expense{}, err
	}

	return e, nil
}

func (r *expenseRepository) Get(ctx context.Context) (expenses []domain.Expense, err error) {
	query := `
		SELECT
			id,
			title,
			amount,
			category,
			created_at
		FROM expenses
		ORDER BY created_at DESC;
	`
	rows, err := r.db.QueryContext(ctx, query) // complete the code
	if err != nil {
		fmt.Printf("repo_error: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e domain.Expense
		if err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.Amount,
			&e.Category,
			&e.CreatedAt,
		); err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *expenseRepository) UpdateById(ctx context.Context, id int64, e *domain.Expense) (expense domain.Expense, err error) {
	query := `
		UPDATE expenses
		SET
			title = $2,
			amount = $3,
			category = $4
		WHERE id = $1
		RETURNING id, title, amount, category, created_at;
	`
	err = r.db.QueryRowContext(ctx, query, id, e.Title, e.Amount, e.Category).Scan(&e.ID, &e.Title, &e.Amount, &e.Category, &e.CreatedAt)
	if err != nil {
		fmt.Printf("repo_error: %s", err.Error())
		return domain.Expense{}, err
	}

	return *e, nil
}

func (r *expenseRepository) DeleteById(ctx context.Context, id int64) error {
	query := `
		DELETE FROM expenses
		WHERE id = $1;
	`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Printf("repo_error: %s", err.Error())
		return err
	}

	return nil
}
