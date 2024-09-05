package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Randomly shuffle a slice and return the first n elements
func randomSubset(slice []string, n int) []string {
	if n > len(slice) {
		n = len(slice)
	}
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice[:n]
}

func main() {
	inputFile := "questions.csv"
	outputFile := "questions_with_company_tags.csv"

	// List of dummy IT company names in lowercase with hyphens for multi-word names
	companies := []string{"google", "microsoft", "meta", "jp-morgan", "amazon", "apple", "ibm", "tesla", "netflix", "adobe"}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Open the input CSV file
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", inputFile, err)
	}
	defer f.Close()

	// Read the CSV file
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Add a new "Company Tags" column to the header
	records[0] = append(records[0], "Company Tags")

	// Iterate over the rows and add random company names
	for i := 1; i < len(records); i++ {
		numCompanies := rand.Intn(len(companies)) + 1 // Random number of companies (at least 1)
		randomCompanies := randomSubset(companies, numCompanies)
		companyTags := strings.Join(randomCompanies, ", ")
		records[i] = append(records[i], companyTags)
	}

	// Create the output CSV file
	outFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", outputFile, err)
	}
	defer outFile.Close()

	// Write the updated records to the output CSV file
	writer := csv.NewWriter(outFile)
	err = writer.WriteAll(records)
	if err != nil {
		log.Fatalf("Failed to write to CSV file: %v", err)
	}

	fmt.Printf("New CSV file with company tags created: %s\n", outputFile)
}
