package pattern

import "testing"

func TestFshuf(t *testing.T) {
	var (
		n     = 10
		pat   = Fshuf{Values: []float32{1, 2, 3}, Repeats: n}
		count = 0
	)
	for _, err := pat.Next(); err == nil; _, err = pat.Next() {
		count++
	}
	if expected, got := 30, count; expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
}
