package main

import (
	"fmt"
	"time"
)

// Disk speed - 1100 Mb/sec

const (
	// Filename - filenate to use
	Filename = "/tmp/testfile"

	// EntriesCount - data set size
	EntriesCount = 5000000

	// AsyncWorkersCount - amount of async workers that write to a file (for async tests)
	AsyncWorkersCount = 100

	// FileBufferSize - size of the buffer for buffered writes
	FileBufferSize = 1024 * 1024

	// PayloadString - payload
	PayloadString = "The idea here is that we are going to open a file and append data to it, tracking what we’re doing with a fixed length: %d\n"
)

// GeneratePayload - prepare data to be written
func GeneratePayload(entriesCount int) (data [][]byte) {
	fmt.Println("Generating payload...")
	start := time.Now()

	data = make([][]byte, entriesCount)

	for i := 0; i < entriesCount; i++ {
		data[i] = []byte(fmt.Sprintf(PayloadString, i))
	}

	fmt.Println("Payload created in", time.Since(start))

	return
}

// RunTest - run N iterations of a test and display average results
func RunTest(proc func(data [][]byte) time.Duration, testName string, data [][]byte, iterations int) {
	fmt.Printf("---> Running %v iterations of test \"%v\"...\n", iterations, testName)

	var elapsedTotal time.Duration

	for i := 0; i < iterations; i++ {
		elapsed := proc(data)
		elapsedTotal += elapsed

		// st, _ := os.Stat(Filename)
		// fmt.Println("Test done in", elapsed, "file size:", st.Size())
	}

	fmt.Printf("Average time on %v iterations: %v\n", iterations, elapsedTotal/time.Duration(iterations))
}

func main() {
	payload := GeneratePayload(EntriesCount)
	iterations := 10

	RunTest(TestFile, "TestFile", payload, iterations)
	RunTest(TestFileAsync, "TestFileAsync", payload, iterations)
	RunTest(TestBufferedFile, "TestBufferedFile", payload, iterations)
	RunTest(TestBufferedFileAsync, "TestBufferedFileAsync", payload, iterations)
	RunTest(TestAsyncFileWithOneWorker, "TestAsyncFileWithOneWorker", payload, iterations)
	RunTest(TestAsyncFileWithManyWorkers, "TestAsyncFileWithManyWorkers", payload, iterations)
	RunTest(TestAsyncFileWithOneWorkerAsync, "TestAsyncFileWithOneWorkerAsync", payload, iterations)
	RunTest(TestAsyncFileWithManyWorkerAsync, "TestAsyncFileWithManyWorkerAsync", payload, iterations)
}
