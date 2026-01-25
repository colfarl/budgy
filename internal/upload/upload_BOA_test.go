package upload

import (
	"errors"
	"testing"
)

func TestIsBoaHeader(t *testing.T) {
	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{input: []string{"Date", "Description", "Amount", "Running Bal."}, want: true},
		{input: []string{"Date", "Description", "Amount"}, want: false},
		{input: []string{"Date", "Description", "Amount", "Random"}, want: false},
		{input: []string{"Date", "Description", "randome", "Running Bal."}, want: false},
		{input: []string{"Date", "random", "Amount", "Running Bal."}, want: false},
		{input: []string{"random", "Description", "Amount", "Running Bal."}, want: false},
	}

	for _, tc := range tests {
		got := isBoaHeaderRow(tc.input)
		if got != tc.want {
			t.Fatalf("Failed: %v\n want: %v got: %v", tc.input, tc.want, got)
		}
	}
}

func TestRowToTransaction(t *testing.T) {
	type test struct {
		row        []string
		account_id int64
		want       transactionParams
		err        error
	}

	tests := []test{
		{
			row:        []string{"11/25/2025", "Beginning balance as of 11/25/2025", "", "1,783.47"},
			account_id: 1,
			want:       transactionParams{},
			err:        ErrInvalidAmount,
		},
		{
			row:        []string{"Date", "Description", "Amount", "Running Bal."},
			account_id: 1,
			want:       transactionParams{},
			err:        ErrInvalidDate,
		},
		{
			row:        []string{"Date", "Description", "Amount"},
			account_id: 1,
			want:       transactionParams{},
			err:        ErrWrongNumCols,
		},
		{

			row:        []string{"11/26/2025", "EXAMPLE TRANSACTION", "-20.00", "0.00"},
			account_id: 1,
			want: transactionParams{
				AccountID:   1,
				Amount:      20.00,
				IsIncome:    false,
				Description: "EXAMPLE TRANSACTION",
				OccurredAt:  "2025-11-26 00:00:00 +0000 UTC",
			},
			err: nil,
		},
		{
			row:        []string{"11/26/2025", "EXAMPLE TRANSACTION", "20.00", "0.00"},
			account_id: 1,
			want: transactionParams{
				AccountID:   1,
				Amount:      20.00,
				IsIncome:    true,
				Description: "EXAMPLE TRANSACTION",
				OccurredAt:  "2025-11-26 00:00:00 +0000 UTC",
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		got, err := rowToTransaction(tc.row, tc.account_id)
		if got != tc.want || !errors.Is(err, tc.err) {
			t.Fatalf("Failed: %v\n want: %v got: %v; err_want: %v, err_got: %v", tc.row, tc.want, got, tc.err, err)
		}
	}
}
