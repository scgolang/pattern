package pattern

// Geom is a geometric series.
type Geom struct {
	Start  float32
	Grow   float32
	Length int
	cur    float32
	idx    int
}

// Next generates the next value in the pattern.
func (pat *Geom) Next() (float32, error) {
	if pat.Length > 0 && pat.idx >= pat.Length {
		return 0, ErrEnd
	}
	if pat.idx == 0 {
		pat.cur = pat.Start
	}
	pat.idx++
	val := pat.cur
	pat.cur = pat.cur * pat.Grow
	return val, nil
}
