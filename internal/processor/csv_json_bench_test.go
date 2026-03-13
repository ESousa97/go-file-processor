package processor

import (
	"encoding/csv"
	"os"
	"testing"
)

func BenchmarkProcess(b *testing.B) {
	// Setup: Generate a temporary CSV file for benchmarking
	source := "bench_input.csv"
	destination := "bench_output.json"
	
	setupBenchData(source, 10000)
	defer os.Remove(source)
	defer os.Remove(destination)

	p := NewCSVToJSONProcessor()

	workerCounts := []int{1, 4, 8, 16}

	for _, wc := range workerCounts {
		// Use fmt or simple string for Run name
		name := "Workers-"
		switch wc {
		case 1: name += "1"
		case 4: name += "4"
		case 8: name += "8"
		case 16: name += "16"
		}
		
		b.Run(name, func(b *testing.B) {
			config := Config{
				WorkerCount: wc,
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := p.Process(source, destination, config)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func setupBenchData(filename string, rows int) {
	f, _ := os.Create(filename)
	w := csv.NewWriter(f)
	w.Write([]string{"id", "name", "email", "role"})
	for i := 0; i < rows; i++ {
		w.Write([]string{"1", "Bench User", "bench@example.com", "worker"})
	}
	w.Flush()
	f.Close()
}
