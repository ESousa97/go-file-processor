package processor

// Processor defines the standard interface for file processing.
// Read reads from a source, Transform converts the data, and Write outputs to a destination.
type Processor interface {
	Process(source, destination string, config Config) (ProcessMetrics, error)
}
