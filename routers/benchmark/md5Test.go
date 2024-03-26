package benchmark

import "crypto/md5"

type TestMd5 struct {
	TestCase
	data []byte
}

func NewTestMd5() *TestMd5 {
	return &TestMd5{
		TestCase: NewTestCase(),
		data:     GenerateData(),
	}
}

func (test *TestMd5) Run(loopCount int) {
	test.StartBenchmarking()
	for i := 0; i < loopCount; i++ {
		md5.Sum(test.data)
	}
	test.StopBenchmarking("Md5")
}
