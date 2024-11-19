package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ReadColumnByName(filename string, columnName string) ([]string, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header: %w", err)
	}

	// Find the index of the specified column
	columnIndex := -1
	for i, name := range header {
		if strings.EqualFold(name, columnName) {
			columnIndex = i
			break
		}
	}
	if columnIndex == -1 {
		return nil, fmt.Errorf("column '%s' not found", columnName)
	}

	// Read the remaining rows and extract the column data
	var columnData []string
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error reading record: %w", err)
		}
		columnData = append(columnData, record[columnIndex])
	}

	return columnData, nil
}
