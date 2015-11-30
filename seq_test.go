package pattern

import "testing"

func TestFseq(t *testing.T) {
	var (
		repeats = 3
		list    = []float32{1, 2, 3, 4, 5}
		pat     = Fseq{Repeats: repeats, Values: list}
	)

	i := 0
	for val, err := pat.Next(); err == nil; val, err = pat.Next() {
		if expected, got := list[i%len(list)], val; expected != got {
			t.Fatalf("Expected %f, got %f", expected, got)
		}
		i++
	}
	if expected, got := 15, i; expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
}

func TestSseq(t *testing.T) {
	var (
		repeats = 3
		list    = []string{"foo", "bar", "baz"}
		pat     = Sseq{Repeats: repeats, Values: list}
	)

	i := 0
	for val, err := pat.Next(); err == nil; val, err = pat.Next() {
		if expected, got := list[i%len(list)], val; expected != got {
			t.Fatalf("Expected %s, got %s", expected, got)
		}
		i++
	}
	if expected, got := 9, i; expected != got {
		t.Fatalf("Expected %d, got %d", expected, got)
	}
}
