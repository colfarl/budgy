package main;

import (
	"fmt"
	"github.com/colfarl/budgy/upload"
)

func printTransactions(transactions [][]string) {
	for _, v := range transactions {
		for _, t := range v {
			fmt.Print(t, " ")	
		}
		fmt.Print("\n")
	}
}

func main() {
	transactions, err := upload.ReadCsvFile("nov-dec.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	printTransactions(transactions)
}
