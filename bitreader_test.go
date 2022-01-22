package bitreader

import "testing"

// TestIncIndex tests that the index will increment or wrap.
func TestIncIndex(t *testing.T) {
	tests := []struct{
		index   int
		want    int
	}{{0, 1}, {7, 0}}

	for _, tt := range tests {
		r := &Reader{
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