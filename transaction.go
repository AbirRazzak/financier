package main

import (
	"log"
	"strconv"
)

type transaction struct {
	date        string // todo convert to date later
	account     string
	description string
	category    string
	tags        string
	amount      float64
}

func newTransaction(record []string) *transaction {
	amountAsString := record[5]
	amountAsFloat, err := strconv.ParseFloat(amountAsString, 32)
	if err != nil {
		log.Fatal(err)
	}

	trans := transaction{
		date:        record[0],
		account:     record[1],
		description: record[2],
		category:    record[3],
		tags:        record[4],
		amount:      amountAsFloat,
	}

	return &trans
}
