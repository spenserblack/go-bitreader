package bitreader

import "testing"

// TestIncIndex tests that the index will increment or wrap.
func TestIncIndex(t *testing.T) {
	tests := []struct{
		bytelen int
		index   int
		want    int
	}{
		{1, 0, 1},
		{1, 7, 0},
		{2, 7, 8},
		{2, 15, 0},
	}

	for _, tt := range tests {
		r := &Reader{
			Bytes: make([]byte, tt.bytelen),
			index: tt.index,
		}
		r.incIndex()
		if r.index != tt.want {
			t.Errorf(
				`Increment from %d with %d bytes = %d, want %d`, 
				tt.index, tt.bytelen, r.index, tt.want,
			)
		}
	}
}