package infra

import (

	"os"
	"errors"
	"time"
	"strconv"
	//"context"
	"encoding/csv"

	//"github.com/colfarl/budgy/internal/database"
	//"github.com/colfarl/budgy/internal/core"
)

type FileProcessor func(string) ([]ParsedTxn, error)

var FileRegistry = map[string]map[string]FileProcessor{
	"boa": {
		"csv": import_boa_csv,  
	},
}

func placeHolder (s string) error {return nil}

var ErrWrongNumCols = errors.New("INVALID TRANSACTION ROW: Wrong number of columns")
var ErrInvalidDate = errors.New("INVALID TRANSACTION ROW: Could not parse date")
var ErrInvalidAmount = errors.New("INVALID TRANSACTION ROW: Could not parse amount")
var ErrNoHeaderRow = errors.New("INVALID TRANSACTION CSV: No header file found")

type ParsedTxn struct {
	Amount		float64
	Description string
	Date		int64	
	Income		bool
}

func import_boa_csv(filename string) ([]ParsedTxn, error) {
	rows, err := ReadCsvFile(filename)
	if err != nil {
		return nil, err
	}

	start_index := -1
	res := make([]ParsedTxn, 0)
	for i := range len(rows) {
		if isBoaHeaderRow(rows[i]) {
			start_index = i
			break
		}
	}

	if start_index == -1 {
		return nil, ErrNoHeaderRow
	}

	for _, v := range rows[start_index+1:] {
		param, err := parseBoaRow(v)
		if err != nil {
			return nil, err
		}
		res = append(res, param)
	}

	return res, nil
}

func parseBoaRow(row []string) (ParsedTxn, error) {
	var parsed_time time.Time
	var parsed_amount float64
	var err error

	if len(row) != 4 {
		return ParsedTxn{}, ErrWrongNumCols
	} else if parsed_time, err = time.Parse("01/02/2006", row[0]); err != nil {
		return ParsedTxn{}, ErrInvalidDate
	} else if parsed_amount, err = strconv.ParseFloat(row[2], 64); err != nil {
		return ParsedTxn{}, ErrInvalidAmount
	} // could check running balance here but we do not use it

	return ParsedTxn{
		Amount:      absFloat64(parsed_amount),
		Date: 		parsed_time.Unix(),
		Description: row[1],
		Income:    parsed_amount > 0,
	}, nil
}


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

func ReadCsvFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}


