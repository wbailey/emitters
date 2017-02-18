package sensors

import "time"

// Sensor
type Sensor interface {
	Measure() float64
	Frequency() time.Duration
}
