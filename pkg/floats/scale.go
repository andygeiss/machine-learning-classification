package floats

// MinMaxScale (Normalization) scales the values down to a value between 0 and 1.
func MinMaxScale(in []float64) {
	min := Minimum(in)
	max := Maximum(in)
	for i := range in {
		in[i] = (in[i] - min) / (max - min)
	}
}

// StandardScale (Standardization) scales the values down to a value between Mean(0) and Standard Deviation (1).
func StandardScale(in []float64) {
	mean := Mean(in)
	stdDev := StandardDeviation(in)
	for i := range in {
		in[i] = (in[i] - mean) / stdDev
	}
}
