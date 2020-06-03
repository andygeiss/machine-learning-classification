package floats_test

import (
	"github.com/andygeiss/machine-learning-classification/pkg/assert"
	"github.com/andygeiss/machine-learning-classification/pkg/floats"
	"testing"
)

func TestVector_Mean_Should_Handle_One_Value(t *testing.T) {
	assert.That(t, floats.Mean([]float64{0}), 0)
}

func TestVector_Mean_Should_Handle_Two_Vector(t *testing.T) {
	assert.That(t, floats.Mean([]float64{0, 2.0}), 1.0)
}

func TestVector_Mean_Should_Handle_Three_Vector(t *testing.T) {
	assert.That(t, floats.Mean([]float64{.5, 1.0, 3.0}), 1.5)
}

func TestVector_Median_Should_Handle_One_Value(t *testing.T) {
	assert.That(t, floats.Median([]float64{1.0}), 1.0)
}

func TestVector_Median_Should_Handle_Two_Vector(t *testing.T) {
	assert.That(t, floats.Median([]float64{.5, 1.0}), .75)
}

func TestVector_Median_Should_Handle_Three_Vector(t *testing.T) {
	assert.That(t, floats.Median([]float64{0, 1.0, 3.0}), 1)
}

func TestVector_Mode_Should_Handle_One_Value(t *testing.T) {
	assert.That(t, floats.Mode([]float64{.5}), .5)
}

func TestVector_Mode_Should_Handle_Two_Vector(t *testing.T) {
	assert.That(t, floats.Mode([]float64{.5, 1.0}), .5)
}

func TestVector_Mode_Should_Handle_Three_Vector(t *testing.T) {
	assert.That(t, floats.Mode([]float64{.5, 1.0, 3.0}), .5)
}
