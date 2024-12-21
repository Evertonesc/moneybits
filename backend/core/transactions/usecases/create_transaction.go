package usecases

import (
	"context"
	"moneybits/core/transactions/domain"
)

type CreateTransactionUseCase struct {
}

func NewCreateTransactionUseCase() *CreateTransactionUseCase {
	return &CreateTransactionUseCase{}
}

// TODO: We must be within a valid planner (be one or another) to link the transaction to it

// A new transaction will always be part of the current month
// every new transaction for a given month must
func (uc *CreateTransactionUseCase) Execute(ctx context.Context, transactionEntry domain.Transaction) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}
