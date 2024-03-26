package benchmark

import "math/rand"

const dataLength = 10000000
const randomSeed = 28

func GenerateData() []byte {
	data := make([]byte, dataLength)
	generator := rand.New(rand.NewSource(randomSeed))
	_, err := generator.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}
