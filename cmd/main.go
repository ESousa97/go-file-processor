package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"go-file-processor/internal/processor"
)

func main() {
	// Setup Structured Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Setup paths
	inputPath := "data/users_large.csv"
	outputPath := "data/output.json"

	// Ensure data directory exists
	if err := os.MkdirAll(filepath.Dir(inputPath), 0755); err != nil {
		slog.Error("Error creating data directory", "error", err)
		os.Exit(1)
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
	slog.Info("Starting file processing", "input", inputPath, "workers", config.WorkerCount)

	metrics, err := csvProcessor.Process(inputPath, outputPath, config)
	if err != nil {
		slog.Error("Critical processing error", "error", err)
		os.Exit(1)
	}

	// Display Execution Summary
	fmt.Printf("\n--- RESUMO DE EXECUÇÃO ---\n")
	fmt.Printf("Total de linhas lidas: %d\n", metrics.TotalLines)
	fmt.Printf("Processadas com sucesso: %d\n", metrics.SuccessCount)
	fmt.Printf("Linhas com erro/ignoradas: %d\n", metrics.ErrorCount)
	fmt.Printf("Tempo total: %s\n", metrics.Duration.Round(time.Millisecond))
	fmt.Printf("--------------------------\n")
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
		slog.Warn("Warning: could not create sample csv", "error", err)
	}
}
