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
			processor.EmailFilter(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`), // Filter invalid emails
			processor.RoleFilter([]string{"administrator", "editor"}),          // Only allow admins and editors
			processor.FieldMasker("role"),                                      // Mask Role field for output
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
		"2,Invalid Email,joe-at-example.com,editor\n" + // Should be filtered out by EmailFilter
		"3,Jane Doe,jane@example.com,viewer\n" +        // Should be filtered out by RoleFilter
		"4,Bob Smith,bob@example.com,editor\n"          // Should stay (but role will be masked)
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Printf("Warning: could not create sample csv: %v", err)
	}
}
