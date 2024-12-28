package dtos

import (
	"moneybits/core/transactions/domain"
	"time"
)

type CreateTransactionRequest struct {
	Amount      int64
	Description string
	Notes       string
	Attachments string // type for files
	Category    string
	Type        string
}

type CreateTransactionResponse struct {
	ID             uint64
	PlannerID      int
	Amount         int64
	Description    string
	Notes          string
	Attachments    string // (files) create a type
	Category       string
	MonthYearIndex string
	Type           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewCreateTXResponseFromDomain(createdTransaction *domain.Transaction) CreateTransactionResponse {
	return CreateTransactionResponse{
		ID:             createdTransaction.ID,
		PlannerID:      createdTransaction.PlannerID,
		Amount:         createdTransaction.Amount,
		Description:    createdTransaction.Description,
		Notes:          createdTransaction.Notes,
		Attachments:    createdTransaction.Attachments,
		Category:       createdTransaction.Category,
		MonthYearIndex: createdTransaction.MonthYearIndex,
		Type:           createdTransaction.Type,
		CreatedAt:      createdTransaction.CreatedAt,
		UpdatedAt:      createdTransaction.UpdatedAt,
	}
}
