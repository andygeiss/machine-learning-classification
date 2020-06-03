package floats_test

import (
	"github.com/andygeiss/machine-learning-classification/pkg/assert"
	"github.com/andygeiss/machine-learning-classification/pkg/floats"
	"testing"
)

func TestKNN_Should_Be_1(t *testing.T) {
	x := []float64{7, 7, 3, 3}
	y := []float64{7, 4, 4, 4}
	z := []float64{0, 0, 1, 1}
	// Replace values which should not be zero with the Mean.
	xMean := floats.Mean(x)
	yMean := floats.Mean(y)
	floats.Replace(x, 0, xMean)
	floats.Replace(y, 0, yMean)
	// Scale the features down to a value between 0 and 1.
	floats.MinMaxScale(x)
	floats.MinMaxScale(x)
	// Finally classify the values
	in := []float64{3, 4}
	floats.MinMaxScale(in)
	observed := floats.Knn(in, x, y, z, 1)
	assert.That(t, observed, 1)
}

func BenchmarkKnn_1000000(b *testing.B) {
	x := randomValues(1000000)
	y := randomValues(1000000)
	z := randomValues(1000000)
	in := []float64{5, 1}
	for i := 0; i < b.N; i++ {
		floats.Knn(in, x, y, z, 1)
	}
}
