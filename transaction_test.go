package main

import (
	"reflect"
	"testing"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransaction(tt.args.record); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
