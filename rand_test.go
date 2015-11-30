package pattern

import "testing"

func TestRand(t *testing.T) {
	var (
		n   = 10
		pat = &Frand{Values: []float32{1, 2, 3, 4}, Length: n}
	)

	for i := 0; i < n; i++ {
		val, err := pat.Next()

		if err != nil && err != End {
			t.Fatal(err)
		}

		if val != 1 && val != 2 && val != 3 && val != 4 {
			t.Fatalf("%f was not one of the provided values", val)
		}
	}
}
