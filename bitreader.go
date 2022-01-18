package bitreader

import "io"

const byteSize int = 8

// Bit is mainly for documentation purposes, and to help clarify the expected
// return value.
type Bit = byte

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

 // ReadBit reads a single bit.
 func (r *Reader) ReadBit() (Bit, error) {
	var err error
	if r.index == 0 {
		_, err = r.R.Read(r.Bytes)
	}
	byteIndex := r.index / byteSize
	bitIndex := r.index % byteSize
	b := r.Bytes[byteIndex] >> (7 - bitIndex)
	b &= 1 // Normalize
	return b, err
 }

// IncIndex increments the index, and loops it to 0 when the max index is
// reached.
func (r *Reader) incIndex() {
	r.index++
	if r.index >= len(r.Bytes) * byteSize {
		r.index = 0
	}
}
