package main

import (
	"bufio"
	"os"
	"sync"
	"time"
)

// AsyncFile - plain unbuffered file
type AsyncFile struct {
	file     *os.File
	buf      *bufio.Writer
	filename string
	channel  chan []byte
	closed   bool
	mutex    sync.Mutex
	workers  int
}

// NewAsyncFile - create new instance of AsyncFile
func NewAsyncFile(filename string, workers int) (file *AsyncFile) {
	file = &AsyncFile{
		filename: filename,
		channel:  make(chan []byte, 1024),
		workers:  workers,
	}

	of, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	file.file = of
	file.buf = bufio.NewWriterSize(of, FileBufferSize)

	for i := 0; i < workers; i++ {
		go file.AsyncFileWriter()
	}

	return
}

// Write - write data to file
func (f *AsyncFile) Write(data []byte) {
	if f.closed {
		return
	}

	f.channel <- data
}

func (f *AsyncFile) internalWrite(data []byte) {
	if f.workers > 1 {
		f.mutex.Lock()
	}

	_, err := f.buf.Write(data)
	if err != nil {
		panic(err)
	}

	if f.workers > 1 {
		f.mutex.Unlock()
	}
}

// AsyncFileWriter - writes data from channel to file buffer
func (f *AsyncFile) AsyncFileWriter() {
	for {
		if f.closed {
			break
		}

		item, ok := <-f.channel
		if !ok {
			break
		}

		f.internalWrite(item)
	}
}

// Close - close file
func (f *AsyncFile) Close() {
	f.closed = true
	close(f.channel)

	for item := range f.channel {
		f.internalWrite(item)
	}

	f.buf.Flush()
	f.file.Sync()
	f.file.Close()
}

// TestAsyncFileWithOneWorker - write data set to AsyncFile and return test performance
func TestAsyncFileWithOneWorker(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewAsyncFile(Filename, 1)

	for i := range data {
		file.Write(data[i])
	}

	elapsed = time.Since(start)
	file.Close()

	return
}

// TestAsyncFileWithManyWorkers - write data set to AsyncFile and return test performance
func TestAsyncFileWithManyWorkers(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewAsyncFile(Filename, 32)

	for i := range data {
		file.Write(data[i])
	}

	elapsed = time.Since(start)
	file.Close()

	return
}

// TestAsyncFileWaitGroup - test AsyncFile with wait group
func TestAsyncFileWaitGroup(data [][]byte, wg *sync.WaitGroup, file *AsyncFile) {
	for i := range data {
		file.Write(data[i])
	}

	wg.Done()
}

// TestAsyncFileWithOneWorkerAsync - test AsyncFile performance
func TestAsyncFileWithOneWorkerAsync(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewAsyncFile(Filename, 1)

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

		go TestAsyncFileWaitGroup(data[begin:end], &wg, file)
	}

	wg.Wait()

	elapsed = time.Since(start)
	file.Close()

	return
}

// TestAsyncFileWithManyWorkerAsync - test AsyncFile performance
func TestAsyncFileWithManyWorkerAsync(data [][]byte) (elapsed time.Duration) {
	os.Remove(Filename)

	start := time.Now()
	file := NewAsyncFile(Filename, 32)

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

		go TestAsyncFileWaitGroup(data[begin:end], &wg, file)
	}

	wg.Wait()

	elapsed = time.Since(start)
	file.Close()

	return
}
