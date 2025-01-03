package usecases

import (
	"context"
	"moneybits/core/transactions/domain"
	"moneybits/core/transactions/dtos"
	"time"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error)
}

type CreateTransactionUseCase struct {
	r TransactionRepository
}

func NewCreateTransactionUseCase(r TransactionRepository) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		r: r,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, createTransactionReq dtos.CreateTransactionRequest) (dtos.CreateTransactionResponse, error) {
	now := time.Now()

	newTransaction, err := domain.NewTransaction(
		createTransactionReq.Amount,
		createTransactionReq.Description,
		createTransactionReq.Notes,
		createTransactionReq.Category,
		createTransactionReq.Type,
		now,
	)
	if err != nil {
		// TODO: create HTTP based errors
		return dtos.CreateTransactionResponse{}, err
	}

	createdTX, err := uc.r.Create(ctx, newTransaction)
	if err != nil {
		return dtos.CreateTransactionResponse{}, err
	}

	return dtos.NewCreateTXResponseFromDomain(createdTX), nil
}
