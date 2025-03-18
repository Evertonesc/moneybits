package domain

import (
	"reflect"
	"testing"
)

func TestCalcSharesTarget(t *testing.T) {
	type args struct {
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
			name: "should return the disired target shares",
			args: args{
				dividendsPaid: []int64{},
				desiredIncome: 1000,
			},
			want:    TargetShares{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcSharesTarget(tt.args.dividendsPaid, tt.args.desiredIncome)
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
