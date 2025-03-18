package domain

import (
	"errors"
	"fmt"
)

type TargetShares struct {
	DividendsPaid []int64
	DesiredIncome int64
	SharesCount   int64
}

func CalcSharesTarget(dividendsPaid []int64, desiredIncome int64) (TargetShares, error) {
	if len(dividendsPaid) == 0 {
		return TargetShares{}, errors.New("a list with dividends paid is required")
	}

	dividendsSum := int64(0)
	for i := 0; i < len(dividendsPaid); i++ {
		dividendsSum += dividendsPaid[i]
	}

	averageDividendsPaid := dividendsSum / int64(len(dividendsPaid))

	fmt.Println(averageDividendsPaid)

	return TargetShares{}, nil
}
