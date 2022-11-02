package main

import (
	"log"
	"strconv"
	"time"
)

type Transaction struct {
	Date        time.Time
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

	dateAsString := record[0]
	dateAsTime, err := time.ParseInLocation("2006-01-02", dateAsString, time.Local) // YYYY-MM-DD
	if err != nil {
		log.Fatal(err)
	}

	trans := Transaction{
		Date:        dateAsTime,
		Account:     record[1],
		Description: record[2],
		Category:    record[3],
		Tags:        record[4],
		Amount:      amountAsFloat,
	}

	return &trans
}
