package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go-file-processor/internal/processor"
)

func main() {
	// Setup paths
	inputPath := "data/users.csv"
	outputPath := "data/output.json"

	// Ensure data directory exists
	if err := os.MkdirAll(filepath.Dir(inputPath), 0755); err != nil {
		log.Fatalf("Error creating data directory: %v", err)
	}

	// Create a dummy CSV if it doesn't exist for test purposes
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		createSampleCSV(inputPath)
	}

	// Initialize Processor
	csvProcessor := processor.NewCSVToJSONProcessor()

	// Execute Pipeline
	fmt.Printf("Processing %s...\n", inputPath)

	if err := csvProcessor.Read(inputPath); err != nil {
		log.Fatalf("Read Error: %v", err)
	}

	if err := csvProcessor.Transform(); err != nil {
		log.Fatalf("Transform Error: %v", err)
	}

	if err := csvProcessor.Write(outputPath); err != nil {
		log.Fatalf("Write Error: %v", err)
	}

	fmt.Printf("Success! Output written to %s\n", outputPath)
}

func createSampleCSV(path string) {
	content := "id,name,email,role\n1,Admin User,admin@example.com,administrator\n2,Regular Joe,joe@example.com,editor\n3,Jane Doe,jane@example.com,viewer\n"
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Printf("Warning: could not create sample csv: %v", err)
	}
}
