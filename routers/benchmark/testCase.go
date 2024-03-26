package benchmark

import (
	"runtime"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
)

type TestCase struct {
	// startCpuTime   time.Duration		// golang cannot get CPU time
	startMemory    uint64
	Result         *TestResult
	stopwatchStart time.Time
}

func NewTestCase() TestCase {
	return TestCase{}
}

func (tc *TestCase) StartBenchmarking() {
	var memStats runtime.MemStats

	runtime.GC()
	debug.SetGCPercent(-1) // disable garbage collector
	runtime.ReadMemStats(&memStats)
	tc.startMemory = memStats.Alloc

	tc.stopwatchStart = time.Now()
}

func (tc *TestCase) StopBenchmarking(algorithm string) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	debug.SetGCPercent(100) // enable garbage collector

	endMemory := memStats.Alloc
	tc.Result = &TestResult{
		ID:            uuid.New().String(),
		Server:        "eric-gin-server",
		Algorithm:     algorithm,
		MemoryUsed:    calcMemoryUsed(tc.startMemory, endMemory),
		ExecutionTime: time.Since(tc.stopwatchStart).Milliseconds(),
		FinishedTime:  time.Now(),
	}
}

func calcMemoryUsed(startMemory uint64, endMemory uint64) int64 {
	if endMemory < startMemory {
		return int64(startMemory-endMemory) * -1
	}
	return int64(endMemory - startMemory)
}
