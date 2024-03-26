package benchmark

import (
	"crypto/md5"
	"crypto/sha256"
	"math/rand"
)

const dataLength = 10000000
const randomSeed = 28

type TestSha256 struct {
	TestCase
	data []byte
}

type TestMd5 struct {
	TestCase
	data []byte
}

func generateData() []byte {
	data := make([]byte, dataLength)
	generator := rand.New(rand.NewSource(randomSeed))
	_, err := generator.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

func NewTestSha256() *TestSha256 {
	return &TestSha256{
		TestCase: NewTestCase(),
		data:     generateData(),
	}
}

func (test *TestSha256) Run(loopCount int) {
	test.StartBenchmarking()
	for i := 0; i < loopCount; i++ {
		sha256.Sum256(test.data)
	}
	test.StopBenchmarking("Sha256")
}

func NewTestMd5() *TestMd5 {
	return &TestMd5{
		TestCase: NewTestCase(),
		data:     generateData(),
	}
}

func (test *TestMd5) Run(loopCount int) {
	test.StartBenchmarking()
	for i := 0; i < loopCount; i++ {
		md5.Sum(test.data)
	}
	test.StopBenchmarking("Md5")
}
