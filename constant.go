package emitters

import "time"

// ConstantEmitter repeats a single value at the given frequency
type ConstantEmitter struct {
	Val  float64
	Freq int //milliseconds
}

// Emit a constant value
func (ce ConstantEmitter) Emit() float64 {
	return ce.Val
}

// Freq reports frequency to emit value
func (ce ConstantEmitter) UseFreq() time.Duration {
	return time.Duration(ce.Freq) * time.Second
}
