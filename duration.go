package pattern

import "time"

// DurGen generates durations.
type DurGen interface {
	Next() (time.Duration, error)
}
