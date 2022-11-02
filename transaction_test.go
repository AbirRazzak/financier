package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_newTransaction(t *testing.T) {
	type args struct {
		record []string
	}
	tests := []struct {
		name string
		args args
		want *Transaction
	}{
		{
			name: "happy path",
			args: args{
				record: []string{"2021-12-01", "Amex Gold Cards", "Paypal", "Other Expenses", "Work, Reimbursable", "50.00"},
			},
			want: &Transaction{
				Date:        time.Date(2021, 12, 01, 0, 0, 0, 0, time.Local),
				Account:     "Amex Gold Cards",
				Description: "Paypal",
				Category:    "Other Expenses",
				Tags:        "Work, Reimbursable",
				Amount:      50.00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransaction(tt.args.record); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
