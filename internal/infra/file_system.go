package infra

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	//"context"
	"encoding/csv"
	//"github.com/colfarl/budgy/internal/database"
	//"github.com/colfarl/budgy/internal/core"
)


type FileKey struct {
	Bank 	string	
	Format	string 	
}

type FileProcessor func(string) ([]ParsedTxn, error)

var FileRegistry = map[FileKey]FileProcessor{
	{Bank: "boa", Format: "csv"}: ImportBoaCsv,
	{Bank: "citi", Format: "csv"}: ImportCitiCsv,
}

var ErrWrongNumCols = errors.New("INVALID TRANSACTION ROW: Wrong number of columns")
var ErrInvalidDate = errors.New("INVALID TRANSACTION ROW: Could not parse date")
var ErrInvalidAmount = errors.New("INVALID TRANSACTION ROW: Could not parse amount")
var ErrNoHeaderRow = errors.New("INVALID TRANSACTION CSV: No header file found")
var ErrUnsupportedFile = errors.New("INVALID FILE: no support for bank and file type")

type ParsedTxn struct {
	Amount		float64
	Description string
	Date		int64	
	Income		bool
}

func ImportCitiCsv(filename string) ([]ParsedTxn, error) {
	rows, err := ReadCsvFile(filename)
	if err != nil {
		return nil, err
	}

	start_index := -1
	res := make([]ParsedTxn, 0)
	for i := range len(rows) {
		if isCitiHeaderRow(rows[i]) {
			start_index = i + 1
			break
		}
	}

	if start_index == -1 {
		return nil, ErrNoHeaderRow
	}

	for _, v := range rows[start_index+1:] {
		log.Printf("%v", v)
		param, err := parseCitiRow(v)
		if err != nil {
			return nil, err
		}
		res = append(res, param)
	}

	return res, nil
}

func ImportBoaCsv(filename string) ([]ParsedTxn, error) {
	rows, err := ReadCsvFile(filename)
	if err != nil {
		return nil, err
	}

	start_index := -1
	res := make([]ParsedTxn, 0)
	for i := range len(rows) {
		if isBoaHeaderRow(rows[i]) {
			start_index = i + 1
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

func parseCitiRow(row []string) (ParsedTxn, error) {
	var parsed_time time.Time
	var parsed_amount float64 
	var cand1 float64
	var cand2 float64
	var err error
	var err2 error

	if len(row) != 6 {
		return ParsedTxn{}, ErrWrongNumCols
	} else if parsed_time, err = time.Parse("01/02/2006", row[1]); err != nil {
		return ParsedTxn{}, ErrInvalidDate
	} 
	
	cand1, err = strconv.ParseFloat(row[3], 64)
	cand2, err2 = strconv.ParseFloat(row[4], 64)
	if cand1 == 0.0 && cand2 == 0.0 {
		return ParsedTxn{}, errors.Join(err, err2) 
	}
	
	if cand1 > 0.0 {
		parsed_amount = cand1
	} else {
		parsed_amount = cand2
	}

	return ParsedTxn{
		Amount:      	absFloat64(parsed_amount),
		Date: 			parsed_time.Unix(),
		Description: 	row[2],
		Income:    		parsed_amount > 0,
	}, nil
}

func parseBoaRow(row []string) (ParsedTxn, error) {
	var parsed_time time.Time
	var parsed_amount float64
	var err error
	
	if len(row) != 4 {
		return ParsedTxn{}, ErrWrongNumCols
	}
	row[2] = strings.ReplaceAll(row[2], ",", "")

	if parsed_time, err = time.Parse("01/02/2006", row[0]); err != nil {
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

func isCitiHeaderRow(csv_row []string) bool{
	if len(csv_row) != 6 {
		return false
	} else if csv_row[0] != "Status" {
		return false
	} else if csv_row[1] != "Date" {
		return false
	} else if csv_row[2] != "Description" {
		return false
	} else if csv_row[3] != "Debit" {
		return false
	} else if csv_row[4] != "Credit" {
		return false
	} else if csv_row[5] != "Member Name" {
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
