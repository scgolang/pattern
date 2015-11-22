package pattern

import "math/rand"

const (
	Inf = 0
)

// Rand emits randomly selected values from an array a
// certain number of times.
func Rand(length int, values []float32) Pattern {
	return func() float32 {
		return values[rand.Intn(len(values))]
	}
}
