package processor

// Config holds the configuration for the file processor, including
// the number of concurrent workers and the chain of data transformers.
type Config struct {
	// WorkerCount specifies the number of parallel goroutines to process data.
	WorkerCount int
	// Transformers is a list of functions applied to each record during processing.
	Transformers []Transformer
}

// AddTransformer appends a new [Transformer] to the configuration chain.
// Transformers are executed sequentially in the order they are added.
func (c *Config) AddTransformer(t Transformer) {
	c.Transformers = append(c.Transformers, t)
}
