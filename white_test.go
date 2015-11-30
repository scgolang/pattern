package pattern

import "testing"

func TestFwhite(t *testing.T) {
	var (
		pat   = &Fwhite{Lo: 2, Hi: 5, Length: 12}
		count = 0
	)

	for val, err := pat.Next(); err == nil; val, err = pat.Next() {
		if val < float32(2) || val >= float32(5) {
			t.Fatalf("%f outside expected range", val)
		}
		count++
	}

	if expected, got := 12, count; expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
}
