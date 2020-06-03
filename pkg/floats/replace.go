package floats

// Replace sets all matching numbers to a new value.
func Replace(in []float64, before, after float64) {
	for i := range in {
		if in[i] == before {
			in[i] = after
		}
	}
}
