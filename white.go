package pattern

import "math/rand"

// Fwhite generates random floats between lo and hi.
type Fwhite struct {
	Lo     float32 `json:"lo"`
	Hi     float32 `json:"hi"`
	Length int     `json:"length"`
	idx    int
}

// Next returns the next value in the pattern.
func (pat *Fwhite) Next() (float32, error) {
	if pat.Length > 0 && pat.idx >= pat.Length {
		return 0, End
	}
	pat.idx++
	return ((pat.Hi - pat.Lo) * rand.Float32()) + pat.Lo, nil
}
