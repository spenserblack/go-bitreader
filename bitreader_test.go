package bitreader

import (
	"bytes"
	"io"
	"testing"
)

// TestIncIndex tests that the index will increment or wrap.
func TestIncIndex(t *testing.T) {
	tests := []struct{
		bytes   int
		index   int
		want    int
	}{
		{1, 0, 1}, {1, 7, 0},
		{2, 0, 1}, {2, 7, 8}, {2, 15, 0},
	}

	for _, tt := range tests {
		r := &Reader{
			bytes: make([]byte, tt.bytes),
			index: tt.index,
		}
		r.incIndex()
		if r.index != tt.want {
			t.Errorf(
				`Increment from %d = %d, want %d`, 
				tt.index, r.index, tt.want,
			)
		}
	}
}

// TestNewBitReader tests that a new bit reader is properly created.
func TestNewBitReader(t *testing.T) {
	r := New(nil, 4)
	if r.r != nil {
		t.Fatalf(`r.r = %v, want nil`, r.r)
	}
	if l := len(r.bytes); l != 4 {
		t.Fatalf(`len(r.bytes) = %v, want 4`, l)
	}
	if r.index != 0 {
		t.Fatalf(`r.index = %v, want 0`, r.index)
	}
}

// TestReadBit tests that a single bit can be read.
func TestReadBit(t *testing.T) {
	buff := bytes.NewBuffer([]byte{0xF0, 0xA5, 0x08})
	tests := []struct {
		want Bit
	}{
		{1}, {1}, {1}, {1}, // F
		{0}, {0}, {0}, {0}, // 0
		{1}, {0}, {1}, {0}, // A
		{0}, {1}, {0}, {1}, // 5
		{0}, {0}, {0}, {0}, // 0
		{1}, {0}, {0}, {0}, // 8
	}
	r := New(buff, 2)

	for i, tt := range tests {
		t.Logf(`bit %d (index %d)`, i, r.index)
		bit, err := r.ReadBit()
		if bit != tt.want {
			t.Errorf(`bit = %d, want %d`, bit, tt.want)
		}
		if err != nil {
			t.Fatalf(`err = %v, want %v`, err, tt.err)
		}
	}
	if _, err := r.ReadBit(); err != io.EOF {
		t.Fatalf(`err = %v, want io.EOF`, err)
	}
}
