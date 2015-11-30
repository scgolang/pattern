package pattern

import "testing"

func TestArith(t *testing.T) {
	pat := &Arith{Start: 3, Step: 4, Length: 5}

	for _, expected := range []float32{3, 7, 11, 15, 19} {
		got, err := pat.Next()
		if err != nil && err != End {
			t.Fatal(err)
		}
		if expected != got {
			t.Fatalf("Expected %f, got %f", expected, got)
		}
	}
}
