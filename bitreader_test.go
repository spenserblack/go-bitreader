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