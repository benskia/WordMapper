package weights

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func BuildWeightMap(filePath string) map[string]int {
	records := extractRecords(filePath)
	return recordsToMap(records)
}

func extractRecords(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error while reading file: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading records: %v\n", err)
		return nil
	}

	return records
}

func recordsToMap(records [][]string) map[string]int {
	mappedWeights := make(map[string]int)
	for _, record := range records {
		keyword := record[0]
		if _, ok := mappedWeights[keyword]; ok {
			fmt.Printf("Tried to map %v - Already mapped", keyword)
		}
		weight, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Printf("Error converting %v to int: %v\n", record[1], err)
		}
		mappedWeights[keyword] = weight
	}
	return mappedWeights
}
