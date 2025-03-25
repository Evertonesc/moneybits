package domain

import (
	"errors"
)

type TargetShares struct {
	Ticker               string
	AverageDividendsPaid int64
	YearsInCalculation   int
	DesiredAnualIncome   int64
	SharesCount          int64
}

type TargetSharesReport struct {
	TargetShares TargetShares
}

// CalcSharesTarget calculates the number of shares needed to reach the desired anual income.
// The calculation is based on a desired anual income.
//
// Using the desired anual income, the function calculates the average of dividends paid in the last years.
// Based on the result, it gets the desired anual income and divides by the average of dividends paid.
// The final result is the number of shares needed to reach the desired anual income.
func CalcSharesTarget(ticker string, dividendsPaid []int64, desiredAnualIncome int64) (TargetShares, error) {
	if len(dividendsPaid) == 0 {
		return TargetShares{}, errors.New("a list with dividends paid is required")
	}

	dividendsSum := int64(0)
	for i := 0; i < len(dividendsPaid); i++ {
		dividendsSum += dividendsPaid[i]
	}

	averageDividendsPaid := dividendsSum / int64(len(dividendsPaid))

	sharesCount := (desiredAnualIncome * 10) / averageDividendsPaid

	return TargetShares{
		Ticker:               ticker,
		SharesCount:          int64(sharesCount),
		DesiredAnualIncome:   18000,
		AverageDividendsPaid: averageDividendsPaid,
		YearsInCalculation:   len(dividendsPaid),
	}, nil
}
