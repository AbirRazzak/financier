package main

import (
	"log"
	"strconv"
)

type Transaction struct {
	Date        string // todo convert to Date later
	Account     string
	Description string
	Category    string
	Tags        string
	Amount      float64
}

func newTransaction(record []string) *Transaction {
	amountAsString := record[5]
	amountAsFloat, err := strconv.ParseFloat(amountAsString, 32)
	if err != nil {
		log.Fatal(err)
	}

	trans := Transaction{
		Date:        record[0],
		Account:     record[1],
		Description: record[2],
		Category:    record[3],
		Tags:        record[4],
		Amount:      amountAsFloat,
	}

	return &trans
}
