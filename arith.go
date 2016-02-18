package pattern

// Arith is a an arithmetic series.
type Arith struct {
	Start  float32 `json:"start"`
	Step   float32 `json:"step"`
	Length int     `json:"length"`
	idx    int
}

// Next returns the next value in the pattern.
func (pat *Arith) Next() (float32, error) {
	if pat.Length > 0 && pat.idx >= pat.Length {
		return 0, ErrEnd
	}
	val := pat.Start + (float32(pat.idx) * pat.Step)
	pat.idx++
	return val, nil
}
