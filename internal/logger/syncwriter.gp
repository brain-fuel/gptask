package logger

import (
	"io"
	"sync"
)

// SynchronizeWriters serializes writes to stdout and stderr through one lock.
// A single lock is required because callers frequently point both streams at
// the same bytes.Buffer or output collector.
func SynchronizeWriters(stdout, stderr io.Writer) (io.Writer, io.Writer) {
	mutex := &sync.Mutex{}
	return &synchronizedWriter{writer: stdout, mutex: mutex},
		&synchronizedWriter{writer: stderr, mutex: mutex}
}

type synchronizedWriter struct {
	writer io.Writer
	mutex  *sync.Mutex
}

func (writer *synchronizedWriter) Write(value []byte) (int, error) {
	writer.mutex.Lock()
	defer writer.mutex.Unlock()
	return writer.writer.Write(value)
}
