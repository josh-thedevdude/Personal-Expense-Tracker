package http

import (
	"encoding/json"
	"net/http"
	"personal_expense_tracker/internal/domain"
	"personal_expense_tracker/internal/service"
	"strconv"
)

type ExpenseHandler struct {
	service service.ExpenseService
}

func NewExpenseHandler(service service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

func (h *ExpenseHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense domain.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateExpense(r.Context(), &expense); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) GetExpenseById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid expense id", http.StatusBadRequest)
		return
	}

	expense, err := h.service.GetExpenseById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.service.GetExpenses(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch expenses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

func (h *ExpenseHandler) UpdateExpenseById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid expense id", http.StatusBadRequest)
		return
	}

	var expense domain.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedExpense, err := h.service.UpdateExpenseById(r.Context(), id, &expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedExpense)
}

func (h *ExpenseHandler) DeleteExpenseById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid expense id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteExpenseById(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
