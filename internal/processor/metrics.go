package processor

import "time"

// ProcessMetrics consolidates the results of the processing.
type ProcessMetrics struct {
	TotalLines   int64
	SuccessCount int64
	ErrorCount   int64
	Duration     time.Duration
}
