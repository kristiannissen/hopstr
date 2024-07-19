package main

import (
	"encoding/csv"
	"log"
	"os"
)

func ImportCSV() {
	file, err := os.Open("./import.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

}
