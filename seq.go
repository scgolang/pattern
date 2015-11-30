package pattern

// Fseq cycles over a list of float values.
type Fseq struct {
	Values  []float32
	Repeats int
	idx     int
	rep     int
}

// Next returns the next value in the pattern.
func (pat *Fseq) Next() (float32, error) {
	if pat.idx >= len(pat.Values) {
		pat.rep++
		if pat.Repeats > 0 && pat.rep >= pat.Repeats {
			return 0, End
		}
		pat.idx = 0
	}
	val := pat.Values[pat.idx]
	pat.idx++
	return val, nil
}

// Sseq cycles over a list of float values.
type Sseq struct {
	Values  []string
	Repeats int
	idx     int
	rep     int
}

// Next returns the next value in the pattern.
func (pat *Sseq) Next() (string, error) {
	if pat.idx >= len(pat.Values) {
		pat.rep++
		if pat.Repeats > 0 && pat.rep >= pat.Repeats {
			return "", End
		}
		pat.idx = 0
	}
	val := pat.Values[pat.idx]
	pat.idx++
	return val, nil
}
