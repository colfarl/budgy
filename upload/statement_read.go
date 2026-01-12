package upload 

import (
	"os"
	"path/filepath"
	"encoding/csv"
)

func getAbsoluteFilepath(filename string) (string, error) {
	currDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	
	path := filepath.Join(currDirectory, "data", filename)
	return path, nil
}

func ReadCsvFile(filename string) ([][]string, error) {
	filepath, err := getAbsoluteFilepath(filename)
	if err != nil {
		return nil, err
	}

    f, err := os.Open(filepath)
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
