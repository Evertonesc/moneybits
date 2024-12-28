package domain

import (
	"errors"
	"fmt"
	"time"
)

const (
	Income  = "INCOME"
	Outcome = "OUTCOME"

	MonthYearTemplate = "Jan-06"
)

type Transaction struct {
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

func NewTransaction(amount int64, description, notes, category, transactionType string, now time.Time) (*Transaction, error) {
	transaction := &Transaction{
		Amount:      amount,
		Description: description,
		Notes:       notes,
		Category:    category,
		Type:        transactionType,
	}

	if !transaction.validTransactionType() {
		return nil, fmt.Errorf("%s is not a valid transaction type", transactionType)
	}

	transaction.ParseMonthYearIndex(now)

	return transaction, nil
}

func (t *Transaction) ParseMonthYearIndex(now time.Time) error {
	if now.IsZero() {
		return errors.New("failed to parse month-year index: the provided date is invalid")
	}

	t.MonthYearIndex = now.Format(MonthYearTemplate)

	return nil
}

func (t *Transaction) validTransactionType() bool {
	return t.Type == Income || t.Type == Outcome
}
