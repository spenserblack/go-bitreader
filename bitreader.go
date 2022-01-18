package bitreader

import "io"

const byteSize int = 8

// Reader reads the bits from a reader that reads bytes.
type Reader struct {
	// R is the underlying reader that reads bytes.
	R      io.Reader
	// Bytes stores the bytes that are read from the reader. When the last
	// available bit is read, the next set of bytes will be read from R. The
	// length of Bytes can be though of as the chunk size.
	Bytes  []byte
	// Index is the index of the bit.
	index  int
}

// IncIndex increments the index, and loops it to 0 when the max index is
// reached.
func (r *Reader) incIndex() {
	r.index++
	if r.index >= len(r.Bytes) * byteSize {
		r.index = 0
	}
}
