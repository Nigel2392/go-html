package gohtml

import (
	"errors"
	"io"
)

// Overflows are not recoverable, and will cause a panic.
var ErrOverflow = errors.New("elementBuffer: buffer overflow")

// A simple buffer for writing HTML elements.
//
// The buffer is not thread-safe, and should not be used concurrently.
//
// The buffer can only be used once, and has a fixed size max size.
type elementBuffer struct {
	// The maximum size of the buffer.
	bufLen int

	// The current index of the buffer.
	//
	// This is the index of the next byte to be written.
	currentIndex int

	// The underlying byte slice.
	buf []byte

	// If true, any overflow will cause a panic.
	panic bool
}

// NewBuffer creates a new elementBuffer with the given size.
//
// If panicOnError is true, any overflow will cause a panic.
func NewBuffer(bufLen int, panicOnError ...bool) *elementBuffer {
	var panicOnErr bool
	if len(panicOnError) > 0 {
		panicOnErr = panicOnError[0]
	}
	return &elementBuffer{
		bufLen:       bufLen,
		buf:          make([]byte, bufLen),
		currentIndex: 0,
		panic:        panicOnErr,
	}
}

// Retrieves the current length of the buffer.
func (b *elementBuffer) Len() int {
	return b.currentIndex
}

// Retrieves the capacity of the buffer.
func (b *elementBuffer) Cap() int {
	return b.bufLen
}

// Retrieves the current buffer as a string.
func (b *elementBuffer) String() string {
	return string(b.buf)
}

// Retrieves the current buffer as a byte slice.
func (b *elementBuffer) Bytes() []byte {
	return b.buf
}

// Writes the given byte slice to the buffer.
func (b *elementBuffer) Write(p []byte) (n int, err error) {
	err = b.overflowCheck(len(p))
	if err != nil {
		return
	}
	copy(b.buf[b.currentIndex:], p)
	b.currentIndex += len(p)
	return len(p), nil
}

// Writes the given string to the buffer.
func (b *elementBuffer) WriteString(s string) (n int, err error) {
	err = b.overflowCheck(len(s))
	if err != nil {
		return
	}
	copy(b.buf[b.currentIndex:], s)
	b.currentIndex += len(s)
	return len(s), nil
}

// Checks if the buffer will overflow if n bytes are written.
func (b *elementBuffer) overflowCheck(n int) error {
	if b.currentIndex+n > b.bufLen {
		if b.panic {
			panic(ErrOverflow)
		}
		return ErrOverflow
	}
	return nil
}

// Writes the buffer to the given writer.
func (b *elementBuffer) WriteTo(w io.Writer) (n int64, err error) {
	var nn int
	nn, err = w.Write(b.buf)
	return int64(nn), err
}
