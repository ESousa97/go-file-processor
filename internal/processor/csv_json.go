package processor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"go-file-processor/internal/domain"
)

// CSVToJSONProcessor implements the Processor interface for CSV to JSON conversion.
type CSVToJSONProcessor struct{}

// NewCSVToJSONProcessor creates a new instance of CSVToJSONProcessor.
func NewCSVToJSONProcessor() *CSVToJSONProcessor {
	return &CSVToJSONProcessor{}
}

// Process executes the full pipeline using a Worker Pool and streaming.
func (p *CSVToJSONProcessor) Process(source, destination string, workerCount int) error {
	// 1. Open source file
	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("could not open source file: %w", err)
	}
	defer srcFile.Close()

	// 2. Open destination file
	dstFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("could not create destination file: %w", err)
	}
	defer dstFile.Close()

	reader := csv.NewReader(srcFile)
	
	// Read headers first
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("could not read csv headers: %w", err)
	}

	headerMap := make(map[string]int)
	for i, h := range headers {
		headerMap[h] = i
	}

	// Channels for Worker Pool
	jobs := make(chan []string, 100)
	results := make(chan domain.User, 100)
	errChan := make(chan error, 1)

	var wg sync.WaitGroup

	// Start Workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range jobs {
				user := domain.User{
					ID:    p.getValue(row, headerMap, "id"),
					Name:  p.getValue(row, headerMap, "name"),
					Email: p.getValue(row, headerMap, "email"),
					Role:  p.getValue(row, headerMap, "role"),
				}
				results <- user
			}
		}()
	}

	// Start Consumer (Writer)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		
		dstFile.WriteString("[\n")
		encoder := json.NewEncoder(dstFile)
		encoder.SetIndent("  ", "  ")
		
		first := true
		for user := range results {
			if !first {
				dstFile.WriteString(",\n")
			}
			encoder.Encode(user)
			first = false
		}
		dstFile.WriteString("\n]")
	}()

	// Producer: Read CSV rows and send to jobs channel
	go func() {
		defer close(jobs)
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				select {
				case errChan <- fmt.Errorf("error reading csv row: %w", err):
				default:
				}
				return
			}
			jobs <- row
		}
	}()

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Wait for consumer to finish
	<-done

	// Check for producer errors
	select {
	case err := <-errChan:
		return err
	default:
	}

	return nil
}

// getValue is a helper to safely retrieve column values by header name.
func (p *CSVToJSONProcessor) getValue(row []string, headerMap map[string]int, key string) string {
	if idx, ok := headerMap[key]; ok && idx < len(row) {
		return row[idx]
	}
	return ""
}
