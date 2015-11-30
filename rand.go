package pattern

import "math/rand"

// Frand creates a stream of random floats.
type Frand struct {
	Values []float32
	Length int
	idx    int
}

// Next returns the next value in the pattern.
func (pat *Frand) Next() (float32, error) {
	if pat.Length > 0 && pat.idx >= pat.Length {
		return 0, End
	}
	pat.idx++
	return pat.Values[rand.Intn(len(pat.Values))], nil
}

// Srand creates a stream of random strings.
type Srand struct {
	Values []string
	Length int
	idx    int
}

// Next returns the next value in the pattern.
func (pat *Srand) Next() (string, error) {
	if pat.Length > 0 && pat.idx >= pat.Length {
		return "", End
	}
	pat.idx++
	return pat.Values[rand.Intn(len(pat.Values))], nil
}
