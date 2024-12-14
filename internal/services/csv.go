package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadColumnByNameAndValidate(filePath string) ([]string, []float64, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	header, err := reader.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("error reading header: %w", err)
	}

	// Find the indices of the AMOUNT and ADDRESS columns
	amountIndex, addressIndex := -1, -1
	for i, name := range header {
		switch strings.ToUpper(strings.TrimSpace(name)) {
		case "AMOUNT":
			amountIndex = i
		case "ADDRESS":
			addressIndex = i
		}
	}
	if amountIndex == -1 || addressIndex == -1 {
		return nil, nil, errors.New("required columns 'AMOUNT' and 'ADDRESS' not found")
	}

	// Regular expression for validating Ethereum addresses
	ethAddressRegex := regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)

	// Read the remaining rows and extract and validate the column data
	var addresses []string
	var amounts []float64

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, nil, fmt.Errorf("error reading record: %w", err)
		}

		// Validate and parse amount
		amountStr := strings.TrimSpace(record[amountIndex])
		if amountStr == "" {
			return nil, nil, errors.New("missing amount in a row")
		}
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil || amount <= 0 {
			return nil, nil, fmt.Errorf("invalid amount '%s': %w", amountStr, err)
		}

		// Validate Ethereum address
		address := strings.TrimSpace(record[addressIndex])
		if address == "" {
			return nil, nil, errors.New("missing address in a row")
		}
		if !ethAddressRegex.MatchString(address) {
			return nil, nil, fmt.Errorf("invalid Ethereum address '%s'", address)
		}

		// Append validated data to respective lists
		amounts = append(amounts, amount)
		addresses = append(addresses, address)
	}

	return addresses, amounts, nil
}
