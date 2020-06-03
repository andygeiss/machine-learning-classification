package floats

import (
	"math"
)

// Maximum returns the greatest number.
func Maximum(in []float64) (out float64) {
	out = float64(0)
	for _, x := range in {
		if x > out {
			out = x
		}
	}
	return
}

// Minimum returns the smallest number.
func Minimum(in []float64) (out float64) {
	out = math.MaxFloat64
	for _, x := range in {
		if x < out {
			out = x
		}
	}
	return
}

// Range measures the absolut difference between the Minimum and Maximum of the numbers.
func Range(in []float64) (out float64) {
	return Maximum(in) - Minimum(in)
}

// StandardDeviation measures how spread out numbers are.
// A higher value means more spread out numbers.
func StandardDeviation(in []float64) (out float64) {
	return math.Sqrt(Variance(in))
}

// Variance measures the average of the squared differences from the Mean.
func Variance(in []float64) (out float64) {
	mean := Mean(in)
	// Sum each difference (to mean) squared.
	for _, x := range in {
		out += (x - mean) * (x - mean)
	}
	// Then average the result.
	return out / float64(len(in))
}
