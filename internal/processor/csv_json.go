package processor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"go-file-processor/internal/domain"
)

// CSVToJSONProcessor implements the Processor interface for CSV to JSON conversion.
type CSVToJSONProcessor struct{}

// NewCSVToJSONProcessor creates a new instance of CSVToJSONProcessor.
func NewCSVToJSONProcessor() *CSVToJSONProcessor {
	return &CSVToJSONProcessor{}
}

// Process executes the full pipeline using a Worker Pool and streaming.
func (p *CSVToJSONProcessor) Process(source, destination string, config Config) (ProcessMetrics, error) {
	start := time.Now()
	metrics := ProcessMetrics{}

	srcFile, dstFile, reader, headerMap, err := p.setupFiles(source, destination)
	if err != nil {
		return metrics, err
	}
	defer srcFile.Close()
	defer dstFile.Close()

	jobs := make(chan []string, 100)
	results := make(chan domain.User, 100)
	done := make(chan bool)

	var wg sync.WaitGroup

	p.startWorkers(&wg, config, jobs, results, headerMap, &metrics)
	go p.runConsumer(results, dstFile, done, &metrics)
	go p.runProducer(reader, jobs, &metrics)

	wg.Wait()
	close(results)
	<-done

	metrics.Duration = time.Since(start)

	// Since we are resilient, we only return error if initialization failed (setupFiles)
	// or if something critical happened that we couldn't log.
	return metrics, nil
}

func (p *CSVToJSONProcessor) setupFiles(source, destination string) (*os.File, *os.File, *csv.Reader, map[string]int, error) {
	srcFile, err := os.Open(source)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("could not open source file: %w", err)
	}

	dstFile, err := os.Create(destination)
	if err != nil {
		srcFile.Close()
		return nil, nil, nil, nil, fmt.Errorf("could not create destination file: %w", err)
	}

	reader := csv.NewReader(srcFile)
	headers, err := reader.Read()
	if err != nil {
		srcFile.Close()
		dstFile.Close()
		return nil, nil, nil, nil, fmt.Errorf("could not read csv headers: %w", err)
	}

	headerMap := make(map[string]int)
	for i, h := range headers {
		headerMap[h] = i
	}

	return srcFile, dstFile, reader, headerMap, nil
}

func (p *CSVToJSONProcessor) startWorkers(wg *sync.WaitGroup, config Config, jobs <-chan []string, results chan<- domain.User, headerMap map[string]int, metrics *ProcessMetrics) {
	for i := 0; i < config.WorkerCount; i++ {
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

				if p.applyTransformers(&user, config.Transformers) {
					results <- user
				}
			}
		}()
	}
}

func (p *CSVToJSONProcessor) applyTransformers(user *domain.User, transformers []Transformer) bool {
	for _, transform := range transformers {
		if !transform(user) {
			return false
		}
	}
	return true
}

func (p *CSVToJSONProcessor) runConsumer(results <-chan domain.User, dstFile *os.File, done chan<- bool, metrics *ProcessMetrics) {
	defer func() { done <- true }()

	dstFile.WriteString("[\n")
	encoder := json.NewEncoder(dstFile)
	encoder.SetIndent("  ", "  ")

	first := true
	for user := range results {
		if !first {
			dstFile.WriteString(",\n")
		}
		if err := encoder.Encode(user); err != nil {
			slog.Error("Failed to encode user to JSON", "error", err, "user_id", user.ID)
			atomic.AddInt64(&metrics.ErrorCount, 1)
			continue
		}
		atomic.AddInt64(&metrics.SuccessCount, 1)
		first = false
	}
	dstFile.WriteString("\n]")
}

func (p *CSVToJSONProcessor) runProducer(reader *csv.Reader, jobs chan<- []string, metrics *ProcessMetrics) {
	defer close(jobs)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		atomic.AddInt64(&metrics.TotalLines, 1)
		if err != nil {
			slog.Warn("Skipping invalid CSV row", "error", err)
			atomic.AddInt64(&metrics.ErrorCount, 1)
			continue
		}
		jobs <- row
	}
}

// getValue is a helper to safely retrieve column values by header name.
func (p *CSVToJSONProcessor) getValue(row []string, headerMap map[string]int, key string) string {
	if idx, ok := headerMap[key]; ok && idx < len(row) {
		return row[idx]
	}
	return ""
}
