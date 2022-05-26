package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var csvFile string

func main() {

	csvFile = os.Args[1]

	lenSlice := len(readCsvFile(csvFile))

	for i := 0; i < lenSlice; i++ {
		records := elSlice(i)
		fmt.Printf("host %s {\nhardware ethernet %s;\nfixed-address  %s;\n}\n", records[2], records[1], records[0])
	}

}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func elSlice(indexSlice int) []string {
	records := readCsvFile(csvFile)
	return records[indexSlice]
}
