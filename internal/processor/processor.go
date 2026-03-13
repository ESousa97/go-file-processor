package processor

// Processor defines the standard interface for file processing tasks.
// Implementations of this interface are responsible for reading data from a source,
// applying transformations, and writing the final result to a destination.
type Processor interface {
	// Process executes the end-to-end transformation workflow and returns [ProcessMetrics].
	// It should be resilient to individual row errors, logging them and continuing
	// the process when possible.
	Process(source, destination string, config Config) (ProcessMetrics, error)
}
