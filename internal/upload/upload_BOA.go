package upload

import (
	"errors"
	"strconv"
	"time"

	"github.com/colfarl/budgy/internal/database"
)

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

func rowToTransaction(csv_row []string, account_id int64) (database.CreateTransactionParams, error) {
	var parsed_time time.Time
	var parsed_amount float64 
	var err error

	if len(csv_row) != 4 {
		return database.CreateTransactionParams{}, errors.New("INVALID TRANSACTION ROW: Wrong number of columns")
	} else if parsed_time, err = time.Parse(time.RFC3339, csv_row[0]); err != nil {
		return database.CreateTransactionParams{}, errors.New("INVALID TRANSACTION ROW: Could not parse date")
	} else if parsed_amount, err = strconv.ParseFloat(csv_row[2], 64); err != nil {
		return database.CreateTransactionParams{}, errors.New("INVALID TRANSACTION ROW: Could not parse amount")
	} // could check running balance here but we do not use it

	return database.CreateTransactionParams{
		AccountID: account_id,	
		OccurredAt: parsed_time.String(), 	
		Amount: parsed_amount,	
		Description: csv_row[1],
		IsIncome: parsed_amount > 0,
	}, nil
}

func parse_boa(csv_rows [][]string) ([]database.CreateTransactionParams, error) {
	start_index := -1
	res := make([]database.CreateTransactionParams, 0)
	for i := range len(csv_rows) {
		if isBoaHeaderRow(csv_rows[i]) {
			start_index = i
			break
		}
	}

	if start_index == -1 {
		return nil, errors.New("INVALID TRANSACTION CSV: No header file found") 
	}
	
	// TODO: Populate res
	for _, _ = range csv_rows[start_index + 1:] {
		continue		
	}

	return res, nil
}

func upload(transactions [][]string, db *database.Queries) error {
	return nil
}
