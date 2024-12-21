package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_ValidTransactionType(t *testing.T) {
	type fields struct {
		ID             uint64
		PlannerID      int
		Amount         int64
		Description    string
		Notes          string
		Attachments    string
		Category       string
		MonthYearIndex string
		Type           string
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid - income transaction type",
			fields: fields{
				ID:             1,
				Type:           Income,
				Amount:         1000,
				Description:    "Salary",
				Category:       "Income",
				MonthYearIndex: "January",
			},
			want: true,
		},
		{
			name: "valid - outcome transaction type",
			fields: fields{
				ID:             2,
				Type:           Outcome,
				Amount:         500,
				Description:    "Groceries",
				Category:       "Food",
				MonthYearIndex: "January",
			},
			want: true,
		},
		{
			name: "invalid - empty transaction type",
			fields: fields{
				ID:             3,
				Type:           "",
				Amount:         100,
				Description:    "Invalid type",
				MonthYearIndex: "January",
			},
			want: false,
		},
		{
			name: "invalid - wrong transaction type",
			fields: fields{
				ID:             4,
				Type:           "INVALID_TYPE",
				Amount:         100,
				Description:    "Wrong type",
				MonthYearIndex: "January",
			},
			want: false,
		},
		{
			name: "invalid - lowercase income type",
			fields: fields{
				ID:             5,
				Type:           "income",
				Amount:         100,
				Description:    "Wrong case",
				MonthYearIndex: "January",
			},
			want: false,
		},
		{
			name: "invalid - lowercase outcome type",
			fields: fields{
				ID:             6,
				Type:           "outcome",
				Amount:         100,
				Description:    "Wrong case",
				MonthYearIndex: "January",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				ID:             tt.fields.ID,
				PlannerID:      tt.fields.PlannerID,
				Amount:         tt.fields.Amount,
				Description:    tt.fields.Description,
				Notes:          tt.fields.Notes,
				Attachments:    tt.fields.Attachments,
				Category:       tt.fields.Category,
				MonthYearIndex: tt.fields.MonthYearIndex,
				Type:           tt.fields.Type,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if got := tr.ValidTransactionType(); got != tt.want {
				t.Errorf("Transaction.ValidTransactionType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_ParseMonthYearIndexYearIndex(t *testing.T) {
	type fields struct {
		ID             uint64
		PlannerID      int
		Amount         int64
		Description    string
		Notes          string
		Attachments    string
		Category       string
		MonthYearIndex string
		Type           string
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantMonthYearIndex string
	}{
		{
			name:               "invalid - empty current date",
			fields:             fields{ID: 1},
			args:               args{now: time.Time{}},
			wantErr:            true,
			wantMonthYearIndex: "",
		},
		{
			name:               "valid - January",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Jan-24",
		},
		{
			name:               "valid - February",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.February, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Feb-24",
		},
		{
			name:               "valid - March",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Mar-24",
		},
		{
			name:               "valid - April",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.April, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Apr-24",
		},
		{
			name:               "valid - May",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.May, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "May-24",
		},
		{
			name:               "valid - June",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Jun-24",
		},
		{
			name:               "valid - July",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.July, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Jul-24",
		},
		{
			name:               "valid - August",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.August, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Aug-24",
		},
		{
			name:               "valid - September",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.September, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Sep-24",
		},
		{
			name:               "valid - October",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.October, 15, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Oct-24",
		},
		{
			name:               "valid - November",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.December, 21, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Dec-24",
		},
		{
			name:               "valid - December",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2024, time.November, 21, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Nov-24",
		},
		{
			name:               "valid - January 2025",
			fields:             fields{ID: 1},
			args:               args{now: time.Date(2025, time.November, 21, 0, 0, 0, 0, time.UTC)},
			wantErr:            false,
			wantMonthYearIndex: "Nov-25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				ID:             tt.fields.ID,
				PlannerID:      tt.fields.PlannerID,
				Amount:         tt.fields.Amount,
				Description:    tt.fields.Description,
				Notes:          tt.fields.Notes,
				Attachments:    tt.fields.Attachments,
				Category:       tt.fields.Category,
				MonthYearIndex: tt.fields.MonthYearIndex,
				Type:           tt.fields.Type,
				CreatedAt:      tt.fields.CreatedAt,
				UpdatedAt:      tt.fields.UpdatedAt,
			}
			if err := tr.ParseMonthYearIndex(tt.args.now); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.ParseMonthYearIndexYearIndex() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !assert.EqualValues(t, tt.wantMonthYearIndex, tr.MonthYearIndex) {
				t.Errorf("month year index mismatch = %s, want %s", tr.MonthYearIndex, tt.wantMonthYearIndex)
			}
		})
	}
}
