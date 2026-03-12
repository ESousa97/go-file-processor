package processor

// Config holds the configuration for the file processor.
type Config struct {
	WorkerCount  int
	Transformers []Transformer
}
