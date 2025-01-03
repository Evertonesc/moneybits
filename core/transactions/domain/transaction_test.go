package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestNewTransaction(t *testing.T) {
	// now := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	type args struct {
		amount          int64
		description     string
		notes           string
		category        string
		transactionType string
		now             time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Transaction
		wantErr bool
	}{
		{
			name: "valid data - income transaction type",
			args: args{
				amount:          1000,
				description:     "Salary",
				notes:           "",
				category:        "Income",
				transactionType: Income,
				now:             time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC),
			},
			want: &Transaction{
				Amount:         1000,
				Description:    "Salary",
				Notes:          "",
				Category:       "Income",
				Type:           Income,
				MonthYearIndex: "Jan-24",
			},
			wantErr: false,
		},
		{
			name: "valid data - outcome transaction type",
			args: args{
				amount:          500,
				description:     "Groceries",
				notes:           "",
				category:        "Food",
				transactionType: Outcome,
				now:             time.Date(2024, time.February, 15, 0, 0, 0, 0, time.UTC),
			},
			want: &Transaction{
				Amount:         500,
				Description:    "Groceries",
				Notes:          "",
				Category:       "Food",
				Type:           Outcome,
				MonthYearIndex: "Feb-24",
			},
			wantErr: false,
		},
		{
			name: "invalid - empty transaction type",
			args: args{
				amount:          100,
				description:     "Invalid type",
				notes:           "",
				category:        "",
				transactionType: "",
				now:             time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid - wrong transaction type",
			args: args{
				amount:          100,
				description:     "Wrong type",
				notes:           "",
				category:        "",
				transactionType: "INVALID_TYPE",
				now:             time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid - lowercase income type",
			args: args{
				amount:          100,
				description:     "Wrong case",
				notes:           "",
				category:        "",
				transactionType: "income",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid - lowercase outcome type",
			args: args{
				amount:          100,
				description:     "Wrong case",
				notes:           "",
				category:        "",
				transactionType: "outcome",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid - empty current datetime",
			args: args{
				amount:          100,
				description:     "Wrong type",
				notes:           "",
				category:        "",
				transactionType: "INVALID_TYPE",
				now:             time.Time{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTransaction(tt.args.amount, tt.args.description, tt.args.notes, tt.args.category, tt.args.transactionType, tt.args.now)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
