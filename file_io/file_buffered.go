package main

import (
	"bufio"
	"os"
	"sync"
	"time"
)

// BufferedFile - plain unbuffered file
type BufferedFile struct {
	file     *os.File
	buf      *bufio.Writer
	mutex    sync.Mutex
	filename string
}

// NewBufferedFile - create new instance of BufferedFile
func NewBufferedFile(filename string) (file *BufferedFile) {
	file = &BufferedFile{
		filename: filename,
	}

	of, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	file.file = of
	file.buf = bufio.NewWriterSize(of, FileBufferSize)

	return
}

// Write - write data to file
func (f *BufferedFile) Write(data []byte) {
	f.mutex.Lock()

	_, err := f.buf.Write(data)
	if err != nil {
		panic(err)
	}

	f.mutex.Unlock()
}

// Close - close file
func (f *BufferedFile) Close() {
	f.buf.Flush()
	f.file.Sync()
	f.file.Close()
}

// TestBufferedFile - write data set to BufferedFile and return test performance
func TestBufferedFile(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewBufferedFile(Filename)

	for i := range data {
		file.Write(data[i])
	}

	elapsed = time.Since(start)
	file.Close()

	return
}

// TestBufferedFileWaitGroup - test BufferedFile with wait group
func TestBufferedFileWaitGroup(data [][]byte, wg *sync.WaitGroup, file *BufferedFile) {
	for i := range data {
		file.Write(data[i])
	}

	wg.Done()
}

// TestBufferedFileAsync - test BufferedFile performance
func TestBufferedFileAsync(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewBufferedFile(Filename)

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

		go TestBufferedFileWaitGroup(data[begin:end], &wg, file)
	}

	wg.Wait()

	elapsed = time.Since(start)
	file.Close()

	return
}
