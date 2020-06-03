package floats

import (
	"math"
	"sort"
)

// Knn ...
func Knn(in []float64, x, y, classifier []float64, k int) float64 {
	// Calculate the distance between the input (in[0], in[1]) and each data point (x[i], y[i]) for i .. n elements in the data set.
	distance := make([]float64, len(x))
	classifierCount := make(map[float64]int)
	classifierMap := make(map[float64]float64, len(x))
	for i := range x {
		// Euclidean Distance
		distance[i] = math.Sqrt(math.Pow(in[0]-x[i], 2) + math.Pow(in[1]-y[i], 2))
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
