package benchmark

import "time"

type TestResult struct {
	ID            string    `json:"id"`
	Server        string    `json:"server"`
	Algorithm     string    `json:"algorithm"`
	MemoryUsed    int64     `json:"memoryUsed"`
	ExecutionTime int64     `json:"executionTime"`
	FinishedTime  time.Time `json:"finishedTime"`
	// CpuTime       time.Duration	// golang cannot get CPU time
}
