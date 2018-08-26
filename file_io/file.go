package main

import (
	"os"
	"sync"
	"time"
)

// File - plain unbuffered file
type File struct {
	file     *os.File
	filename string
}

// NewFile - create new instance of File
func NewFile(filename string) (file *File) {
	file = &File{
		filename: filename,
	}

	of, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	file.file = of

	return
}

// Write - write data to file
func (f *File) Write(data []byte) {
	f.file.Write(data)
}

// Close - close file
func (f *File) Close() {
	f.file.Sync()
	f.file.Close()
}

// TestFile - write data set to File and return test performance
func TestFile(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewFile(Filename)

	for i := range data {
		file.Write(data[i])
	}

	elapsed = time.Since(start)
	file.Close()

	return
}

// TestFileWaitGroup - test File with wait group
func TestFileWaitGroup(data [][]byte, wg *sync.WaitGroup, file *File) {
	for i := range data {
		file.Write(data[i])
	}

	wg.Done()
}

// TestFileAsync - test File performance
func TestFileAsync(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewFile(Filename)

	var wg sync.WaitGroup
	wg.Add(AsyncWorkersCount)

	for i := 0; i < AsyncWorkersCount; i++ {
		batchSize := len(data) / AsyncWorkersCount
		begin := i * batchSize
		end := (i + 1) * batchSize

		remaining := len(data) - end
		if remaining > 0 && remaining < batchSize {
			end += remaining
		}

		go TestFileWaitGroup(data[begin:end], &wg, file)
	}

	wg.Wait()

	elapsed = time.Since(start)
	file.Close()

	return
}
