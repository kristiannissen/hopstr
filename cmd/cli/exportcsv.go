package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func ExportCSV() {
	var b []byte
	var err error
	// Readfile file
	if b, err = os.ReadFile("./hops.json"); err != nil {
		log.Println(err)
	}
	// unmarshal
	if err = json.Unmarshal(b, &Hops); err != nil {
		log.Println(err)
	}

	// Create CSV file
	file, ferr := os.Create("./export.csv")
	if ferr != nil {
		log.Println(ferr)
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Iterate over hops
	for _, hop := range Hops {
		// Extract data
		for _, data := range hop.Data {
			if strings.HasPrefix(data.Name, "Flavor") {
				// Remove the unwanted part
				indx := strings.LastIndex(data.Value, "These are the common")
				writer.Write([]string{hop.Uuid, hop.Name, data.Value[0:indx]})
			}
		}

	}
}
