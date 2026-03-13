package processor

// Config holds the configuration for the file processor.
type Config struct {
	WorkerCount  int
	Transformers []Transformer
}

// AddTransformer adds a new transformer to the configuration chain.
func (c *Config) AddTransformer(t Transformer) {
	c.Transformers = append(c.Transformers, t)
}
