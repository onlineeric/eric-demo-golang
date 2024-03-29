package benchmark

import "crypto/sha256"

type TestSha256 struct {
	TestCase
	data []byte
}

func NewTestSha256() *TestSha256 {
	return &TestSha256{
		TestCase: NewTestCase(),
		data:     GenerateData(),
	}
}

func (test *TestSha256) Run(loopCount int) {
	test.StartBenchmarking()
	for i := 0; i < loopCount; i++ {
		sha256.Sum256(test.data)
	}
	test.StopBenchmarking("SHA256")
}
