package pattern

// F generates a constant float
type F float32

// Next returns the next value in the pattern.
func (pat F) Next() (float32, error) {
	return float32(pat), nil
}
