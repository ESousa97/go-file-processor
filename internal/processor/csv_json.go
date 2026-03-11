package processor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"go-file-processor/internal/domain"
)

// CSVToJSONProcessor implements the Processor interface for CSV to JSON conversion.
type CSVToJSONProcessor struct {
	sourceData [][]string
	records    []domain.User
}

// NewCSVToJSONProcessor creates a new instance of CSVToJSONProcessor.
func NewCSVToJSONProcessor() *CSVToJSONProcessor {
	return &CSVToJSONProcessor{}
}

// Read opens and reads the CSV file content.
func (p *CSVToJSONProcessor) Read(source string) error {
	file, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("could not open source file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read csv data: %w", err)
	}

	if len(data) < 2 {
		return fmt.Errorf("csv file is empty or missing headers")
	}

	p.sourceData = data
	return nil
}

// Transform maps the CSV columns to the User struct.
// It assumes the first row contains headers matching the struct field names.
func (p *CSVToJSONProcessor) Transform() error {
	headers := p.sourceData[0]
	rows := p.sourceData[1:]

	headerMap := make(map[string]int)
	for i, h := range headers {
		headerMap[h] = i
	}

	p.records = make([]domain.User, 0, len(rows))

	for _, row := range rows {
		user := domain.User{
			ID:    p.getValue(row, headerMap, "id"),
			Name:  p.getValue(row, headerMap, "name"),
			Email: p.getValue(row, headerMap, "email"),
			Role:  p.getValue(row, headerMap, "role"),
		}
		p.records = append(p.records, user)
	}

	return nil
}

// Write encodes the records into JSON and writes to the destination path.
func (p *CSVToJSONProcessor) Write(destination string) error {
	file, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("could not create destination file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(p.records); err != nil {
		return fmt.Errorf("could not encode records to json: %w", err)
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
