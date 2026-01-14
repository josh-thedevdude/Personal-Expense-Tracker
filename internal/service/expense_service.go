package service

import (
	"context"
	"errors"
	"personal_expense_tracker/internal/domain"
	"personal_expense_tracker/internal/repository"
)

// ExpenseService defines business-level operations
type ExpenseService interface {
	CreateExpense(ctx context.Context, e *domain.Expense) error
	GetExpenseById(ctx context.Context, id int64) (domain.Expense, error)
	GetExpenses(ctx context.Context) ([]domain.Expense, error)
	UpdateExpenseById(ctx context.Context, id int64, e *domain.Expense) (domain.Expense, error)
	DeleteExpenseById(ctx context.Context, id int64) error
}

type expenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{
		repo: repo,
	}
}

func (s *expenseService) CreateExpense(ctx context.Context, e *domain.Expense) error {
	if e.Title == "" {
		return errors.New("title cannot be empty")
	}

	if e.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return s.repo.Create(ctx, e)
}

func (s *expenseService) GetExpenseById(ctx context.Context, id int64) (domain.Expense, error) {
	if id <= 0 {
		return domain.Expense{}, errors.New("invalid expense id")
	}

	return s.repo.GetById(ctx, id)
}

func (s *expenseService) GetExpenses(ctx context.Context) ([]domain.Expense, error) {
	return s.repo.Get(ctx)
}

func (s *expenseService) UpdateExpenseById(
	ctx context.Context,
	id int64,
	e *domain.Expense,
) (domain.Expense, error) {

	if id <= 0 {
		return domain.Expense{}, errors.New("invalid expense id")
	}

	if e.Title == "" {
		return domain.Expense{}, errors.New("title cannot be empty")
	}

	if e.Amount <= 0 {
		return domain.Expense{}, errors.New("amount must be greater than zero")
	}

	return s.repo.UpdateById(ctx, id, e)
}

func (s *expenseService) DeleteExpenseById(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid expense id")
	}

	return s.repo.DeleteById(ctx, id)
}
