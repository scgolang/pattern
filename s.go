package pattern

// S returns a string on every invocation of Next.
type S string

// Next returns the next value in the pattern.
func (s S) Next() (string, error) {
	return string(s), nil
}
