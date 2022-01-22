package bitreader

import "io"

const byteSize int = 8

// Bit is mainly for documentation purposes, and to help clarify the expected
// return value.
type Bit = byte

// Reader reads the bits from a reader that reads bytes.
type Reader struct {
	// R is the underlying reader that reads bytes.
	r io.Reader
	// Bytes stores the bytes that are read from the reader. When the last
	// available bit is read, the next set of bytes will be read from R. The
	// length of Bytes can be though of as the chunk size.
	bytes []byte
	// Index is the index of the bit.
	index int
}

// New creates a new bit reader. The amount of bytes to be read at a time is
// set by chunkSize.
func New(r io.Reader, chunkSize int) *Reader {
	return &Reader{
		r: r,
		bytes: make([]byte, chunkSize),
	}
}

 // ReadBit reads a single bit.
 func (r *Reader) ReadBit() (Bit, error) {
	var err error
	if r.index == 0 {
		_, err = r.r.Read(r.bytes)
	}
	b := r.currentBit()
	r.incIndex()
	return b, err
 }

// IncIndex increments the index, and loops it to 0 when the max index is
// reached.
func (r *Reader) incIndex() {
	r.index++
	r.index %= len(r.bytes)
}

// ByteIndex gets the index of the byte that the bit is read from.
func (r Reader) byteIndex() int {
	return r.index / byteSize
}

// BitIndex gets the index of the bit within the byte (counting from the left).
func (r Reader) bitIndex() int {
	return r.index % byteSize
}

// CurrentByte gets the current byte to read.
func (r Reader) currentByte() byte {
	return r.bytes[r.byteIndex()]
}

// CurrentBit gets the current bit to read.
func (r Reader) currentBit() Bit {
	bitIndex := 7 - r.bitIndex()
	return (r.currentByte() >> bitIndex) & 1
}
