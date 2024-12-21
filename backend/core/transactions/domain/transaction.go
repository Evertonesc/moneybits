package domain

import (
	"errors"
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
	Attachments    string // create a type
	Category       string
	MonthYearIndex string
	Type           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (t *Transaction) ParseMonthYearIndex(now time.Time) error {
	if now.IsZero() {
		return errors.New("failed to parse month-year index: the provided date is invalid")
	}

	t.MonthYearIndex = now.Format(MonthYearTemplate)

	return nil
}

func (t *Transaction) ValidTransactionType() bool {
	return t.Type == Income || t.Type == Outcome
}
