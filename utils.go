package utils

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// RWio interface shows
// available operations
// on read and write
type RWio interface {
	Read() string
	Write(string, int) (int, error)
}

// RW is a structure for reader
// and writer
type RW struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

// Read implements RWio.
func (r *RW) Read() string {
	readData, err := r.reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading data")
	}
	return strings.TrimSpace(readData)
}

// Write implements RWio.
func (w *RW) Write(s string, bufferSize int) (int, error) {
	fmt.Println("writing data")
	len, err := w.writer.WriteString(s)
	if err != nil {
		log.Println("flusing the buffer")
		w.writer.Flush()
		log.Println("flushed the buffer available for new write")
		return len, nil
	}
	return len, nil

}

// WithWriter returns a writer initialized RW struct
func (w *RW) WithWriter(writer *bufio.Writer) *RW {
	w.writer = writer
	return w

}

// WithReader returns a reader initialized RW struct
func (r *RW) WithReader(reader *bufio.Reader) *RW {
	r.reader = reader
	return r
}

// NewRW initializes both reader and writer and returns
func NewRW(writer *bufio.Writer, reader *bufio.Reader) RWio {
	return &RW{
		reader: reader,
		writer: writer,
	}
}
