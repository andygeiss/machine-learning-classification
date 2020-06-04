package floats

import (
	"math"
	"sort"
)

// ManhattanDistance ...
func ManhattanDistance(x1, x2, y1, y2 float64) float64 {
	return math.Abs(x1-x2) + math.Abs(y1-y2)
}

// EuclideanDistance ...
func EuclideanDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

// Knn ...
func Knn(in []float64, x, y, classifier []float64, k int, distanceFn func(x1, x2, y1, y2 float64) float64) float64 {
	// Calculate the distance between the input (in[0], in[1]) and each data point (x[i], y[i]) for i .. n elements in the data set.
	distance := make([]float64, len(x))
	classifierCount := make(map[float64]int)
	classifierMap := make(map[float64]float64, len(x))
	for i := range x {
		// Calculate the distance
		distance[i] = distanceFn(in[0], x[i], in[1], y[i])
		// Map the distance value to the class identifier classifier
		classifierMap[distance[i]] = classifier[i]
	}
	// Sort the distance to pick up the first k elements later.
	sort.Float64s(distance)
	// Look at the first k elements count the votes for each class identifier.
	// Thus leads to ensure an odd k value, because with an even value it is possible to get more than one class.
	for i := 0; i < k; i++ {
		classifierCount[classifierMap[distance[i]]]++
	}
	// Finally pick the class with the most votes.
	max := 0
	out := float64(0)
	for k, v := range classifierCount {
		if v > max {
			out = k
		}
	}
	return out
}
