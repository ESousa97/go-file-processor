package processor

import "time"

// ProcessMetrics consolidates the statistical results of a file processing operation.
type ProcessMetrics struct {
	// TotalLines is the total number of lines read from the source file.
	TotalLines int64
	// SuccessCount is the number of records successfully processed and written.
	SuccessCount int64
	// ErrorCount is the number of records that failed validation or transformation.
	ErrorCount int64
	// Duration is the total time taken to complete the processing task.
	Duration time.Duration
}
