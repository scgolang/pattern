package pattern

import "errors"

var (
	// ErrEnd is the error returned by a pattern to
	// signal that the pattern has ended.
	ErrEnd = errors.New("end of pattern")
)

// FloatGen is anything that generates floats
// with a Next method.
type FloatGen interface {
	Next() (float32, error)
}

// StringGen is anything that generates strings
// with a Next method.
type StringGen interface {
	Next() (string, error)
}
