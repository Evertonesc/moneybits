package dtos

import (
	"moneybits/core/transactions/domain"
	"time"
)

type CreateTransactionRequest struct {
	Amount      int64  `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Notes       string `json:"notes,omitempty"`
	Attachments string `json:"attachments,omitempty"` // type for files
	Category    string `json:"category,omitempty"`
	Type        string `json:"type,omitempty"`
}

type CreateTransactionResponse struct {
	ID             uint64    `json:"id,omitempty"`
	Amount         int64     `json:"amount,omitempty"`
	Description    string    `json:"description,omitempty"`
	Notes          string    `json:"notes,omitempty"`
	Attachments    string    `json:"attachments,omitempty"` // (files) create a type
	Category       string    `json:"category,omitempty"`
	MonthYearIndex string    `json:"month_year_index,omitempty"`
	Type           string    `json:"type,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func NewCreateTXResponseFromDomain(createdTransaction *domain.Transaction) CreateTransactionResponse {
	return CreateTransactionResponse{
		ID:             createdTransaction.ID,
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
