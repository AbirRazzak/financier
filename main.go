package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func readTransactions(file string) <-chan []string {
	//f, err := os.Open(file)
	f, _ := os.Open(file)
	defer f.Close()
	//if err != nil {
	//	return "", err
	//}

	r := csv.NewReader(f)

	// skip the first line
	r.Read()
	//if _, err := r.Read(); err != nil {
	//	return "", err
	//}

	out := make(chan []string)
	go func() {
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			out <- record
			time.Sleep(100000000)
		}
		close(out)
	}()
	return out
}

func convertTransactionToStructure(in <-chan []string) <-chan *Transaction {
	out := make(chan *Transaction)
	go func() {
		for record := range in {
			out <- newTransaction(record)
		}
		close(out)
	}()

	return out
}

func main() {
	readFilesChannel := readTransactions("C:\\Users\\Abir\\Downloads\\transactions.csv")
	convertRawDataChannel := convertTransactionToStructure(readFilesChannel)

	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
	fmt.Println(<-convertRawDataChannel)
}
