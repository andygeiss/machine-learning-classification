package floats

import "sort"

// Mean is the average of the given values.
func Mean(in []float64) (out float64) {
	var sum float64
	for _, x := range in {
		sum += x
	}
	return sum / float64(len(in))
}

// Median is the number that separates the lower half of the values from the upper half.
func Median(in []float64) (out float64) {
	sort.Float64s(in)
	if len(in)%2 == 0 {
		return (in[len(in)/2] + in[len(in)/2-1]) / 2
	}
	return in[len(in)/2]
}

// Mode is the most repeated number.
func Mode(in []float64) (out float64) {
	var count, max, val float64
	for _, cur := range in {
		count = 0
		for _, y := range in {
			if y == cur {
				count++
			}
		}
		if count > max {
			max = count
			val = cur
		}
	}
	return val
}
