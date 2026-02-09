package upload

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/colfarl/budgy/internal/database"
)

type querier = database.Queries
type transactionParams = database.CreateTransactionParams

var ErrWrongNumCols = errors.New("INVALID TRANSACTION ROW: Wrong number of columns")
var ErrInvalidDate = errors.New("INVALID TRANSACTION ROW: Could not parse date")
var ErrInvalidAmount = errors.New("INVALID TRANSACTION ROW: Could not parse amount")
var ErrNoHeaderRow = errors.New("INVALID TRANSACTION CSV: No header file found")

const (
	DateLayoutBOA = "01/02/2006"
)

// TODO: maybe move later
func absFloat64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func isBoaHeaderRow(csv_row []string) bool {
	if len(csv_row) != 4 {
		return false
	} else if csv_row[0] != "Date" {
		return false
	} else if csv_row[1] != "Description" {
		return false
	} else if csv_row[2] != "Amount" {
		return false
	} else if csv_row[3] != "Running Bal." {
		return false
	}
	return true
}

func rowToTransaction(csv_row []string, account_id int64) (transactionParams, error) {
	var parsed_time time.Time
	var parsed_amount float64
	var err error

	if len(csv_row) != 4 {
		return transactionParams{}, ErrWrongNumCols
	} else if parsed_time, err = time.Parse(DateLayoutBOA, csv_row[0]); err != nil {
		return transactionParams{}, ErrInvalidDate
	} else if parsed_amount, err = strconv.ParseFloat(csv_row[2], 64); err != nil {
		return transactionParams{}, ErrInvalidAmount
	} // could check running balance here but we do not use it

	return transactionParams{
		AccountID:   account_id,
		OccurredAt:  parsed_time.String(),
		Amount:      absFloat64(parsed_amount),
		Description: csv_row[1],
		IsIncome:    parsed_amount > 0,
	}, nil
}

func parse_boa(csv_rows [][]string, account_id int64) ([]transactionParams, error) {
	start_index := -1
	res := make([]transactionParams, 0)
	for i := range len(csv_rows) {
		if isBoaHeaderRow(csv_rows[i]) {
			start_index = i
			break
		}
	}

	if start_index == -1 {
		return nil, ErrNoHeaderRow
	}

	for _, v := range csv_rows[start_index+1:] {
		param, err := rowToTransaction(v, account_id)
		if err != nil {
			return nil, err
		}
		res = append(res, param)
	}

	return res, nil
}
