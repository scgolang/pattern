package pattern

import "math/rand"

// Fshuf generates a pattern that is a pseudo-random
// permutation of the list of values.
type Fshuf struct {
	Values  []float32 `json:"values"`
	Repeats int       `json:"repeats"`
	indices []int
	idx     int
	rep     int
}

// Next returns the next value in the pattern.
func (pat *Fshuf) Next() (float32, error) {
	if pat.idx == 0 && pat.rep == 0 {
		pat.indices = rand.Perm(len(pat.Values))
	}
	if pat.idx >= len(pat.Values) {
		pat.rep++
		if pat.Repeats > 0 && pat.rep >= pat.Repeats {
			return 0, End
		}
		pat.idx = 0
	}
	val := pat.Values[pat.indices[pat.idx]]
	pat.idx++
	return val, nil
}

// Sshuf generates a pattern that is a pseudo-random
// permutation of the list of values.
type Sshuf struct {
	Values  []string `json:"values"`
	Repeats int      `json:"repeats"`
	indices []int
	idx     int
	rep     int
}

// Next returns the next value in the pattern.
func (pat *Sshuf) Next() (string, error) {
	if pat.idx == 0 && pat.rep == 0 {
		pat.indices = rand.Perm(len(pat.Values))
	}
	if pat.idx >= len(pat.Values) {
		pat.rep++
		if pat.Repeats > 0 && pat.rep >= pat.Repeats {
			return "", End
		}
		pat.idx = 0
	}
	val := pat.Values[pat.indices[pat.idx]]
	pat.idx++
	return val, nil
}
