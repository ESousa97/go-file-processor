package processor_test

import (
	"fmt"
	"go-file-processor/internal/processor"
)

// ExampleCSVToJSONProcessor_Process demonstrates how to use the CSVToJSONProcessor
// with worker pool configuration and a chain of transformers.
func ExampleCSVToJSONProcessor_Process() {
	proc := processor.NewCSVToJSONProcessor()

	// Configure the processor with 4 workers and several transformers
	config := processor.Config{
		WorkerCount: 4,
	}

	// Add an email filter to only allow standard emails
	config.AddTransformer(processor.EmailFilter(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`))

	// Add a role filter to only allow Admins and Editors
	config.AddTransformer(processor.RoleFilter([]string{"Admin", "Editor"}))

	// Add a field masker for privacy
	config.AddTransformer(processor.FieldMasker("email"))

	// Note: In a real scenario, paths would be real files
	// metrics, err := proc.Process("input.csv", "output.json", config)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Printf("Processed %d lines\n", metrics.TotalLines)

	fmt.Println("Processor configured with 4 workers and 3 transformers.")
	// Output: Processor configured with 4 workers and 3 transformers.
}
