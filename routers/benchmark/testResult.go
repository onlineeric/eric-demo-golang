package benchmark

import "time"

type TestResult struct {
	// CpuTime       time.Duration	// golang cannot get CPU time
	MemoryUsed    int64
	ExecutionTime int64
	FinishedTime  time.Time
}
