package emitters

import "time"

// Emitter
type Emitter interface {
	Emit() float64
	Freq() time.Duration
}
