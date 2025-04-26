package domain

import (
	"reflect"
	"testing"
)

func TestCalcSharesTarget(t *testing.T) {
	type args struct {
		ticker        string
		dividendsPaid []int64
		desiredIncome int64
	}
	tests := []struct {
		name    string
		args    args
		want    TargetShares
		wantErr bool
	}{
		{
			name: "should return error when dividends paid is nil",
			args: args{
				dividendsPaid: nil,
				desiredIncome: 1000,
			},
			want:    TargetShares{},
			wantErr: true,
		},
		{
			name: "should return error when dividends paid is empty",
			args: args{
				dividendsPaid: []int64{},
				desiredIncome: 1000,
			},
			want:    TargetShares{},
			wantErr: true,
		},
		{
			name: "should return the disired target shares for TAEE4",
			args: args{
				ticker: "TAEE4",
				dividendsPaid: []int64{
					118,
					97,
					162,
					150,
					107,
					63,
				},
				desiredIncome: 180000,
			},
			want: TargetShares{
				Ticker:               "TAEE4",
				SharesCount:          int64(1551),
				DesiredAnualIncome:   180000,
				AverageDividendsPaid: 116,
				YearsInCalculation:   6,
			},
			wantErr: false,
		},
		{
			name: "should return the disired target shares for BBAS3",
			args: args{
				ticker: "BBAS3",
				dividendsPaid: []int64{
					260,
					227,
					207,
					113,
				},
				desiredIncome: 180000,
			},
			want: TargetShares{
				Ticker:               "BBAS3",
				SharesCount:          int64(895),
				DesiredAnualIncome:   180000,
				AverageDividendsPaid: 201,
				YearsInCalculation:   4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcSharesTarget(tt.args.ticker, tt.args.dividendsPaid, tt.args.desiredIncome)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcSharesTarget() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcSharesTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
