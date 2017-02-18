package sensors

import "time"

// ConstantSensor repeats a single value at the given frequency
type ConstantSensor struct {
	Val  float64
	Freq int //milliseconds
}

// Measure a constant value
func (ce ConstantSensor) Measure() float64 {
	return ce.Val
}

// Frequency to measure the value
func (ce ConstantSensor) Frequency() time.Duration {
	return time.Duration(ce.Freq) * time.Second
}
