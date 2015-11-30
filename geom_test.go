package pattern

import "testing"

func TestGeom(t *testing.T) {
	pat := &Geom{Start: 1, Grow: 2, Length: 5}

	for _, expected := range []float32{1, 2, 4, 8, 16} {
		got, err := pat.Next()
		if err != nil && err != End {
			t.Fatal(err)
		}
		if expected != got {
			t.Fatalf("Expected %f, got %f", expected, got)
		}
	}
}
