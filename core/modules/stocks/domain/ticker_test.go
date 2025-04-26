package domain

import (
	"reflect"
	"testing"
)

func TestNewTicker(t *testing.T) {
	type args struct {
		ticker string
	}
	tests := []struct {
		name    string
		args    args
		want    Ticker
		wantErr bool
	}{
		{
			name:    "Empty ticker should return error",
			args:    args{ticker: ""},
			want:    Ticker{},
			wantErr: true,
		},
		{
			name:    "Invalid ticker format should return error",
			args:    args{ticker: "ThisIsAReallyLongRandomStringThatIsDefinitelyNotAValidStockTickerSymbolAtAllAndHasMoreThan60Characters"},
			want:    Ticker{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTicker(tt.args.ticker)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTicker() = %v, want %v", got, tt.want)
			}
		})
	}
}
