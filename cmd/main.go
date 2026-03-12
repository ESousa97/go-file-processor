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
	inputPath := "data/users_large.csv"
	outputPath := "data/output.json"

	// Ensure data directory exists
	if err := os.MkdirAll(filepath.Dir(inputPath), 0755); err != nil {
		log.Fatalf("Error creating data directory: %v", err)
	}

	// Create a dummy CSV if it doesn't exist for test purposes
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		createSampleCSV(inputPath)
	}

	// Initialize Processor with Config
	config := processor.Config{
		WorkerCount: 5,
		Transformers: []processor.Transformer{
			processor.EmailValidatorTransformer(),    // Filter invalid emails
			processor.SensitiveDataTransformer(true), // Mask Role field
		},
	}
	csvProcessor := processor.NewCSVToJSONProcessor()

	// Execute Pipeline
	fmt.Printf("Processing %s with %d workers and transformations...\n", inputPath, config.WorkerCount)

	if err := csvProcessor.Process(inputPath, outputPath, config); err != nil {
		log.Fatalf("Processing Error: %v", err)
	}

	fmt.Printf("Success! Output written to %s\n", outputPath)
}

func createSampleCSV(path string) {
	// Including some invalid data to test transformers
	content := "id,name,email,role\n" +
		"1,Admin User,admin@example.com,administrator\n" +
		"2,Invalid Email,joe-at-example.com,editor\n" + // Should be filtered out
		"3,Jane Doe,jane@example.com,viewer\n"
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Printf("Warning: could not create sample csv: %v", err)
	}
}
