package emitters

import "time"

// Emitter
type Emitter interface {
	Emit() float64
	UseFreq() time.Duration
}
