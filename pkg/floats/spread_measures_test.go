package floats_test

import (
	"github.com/andygeiss/machine-learning-classification/pkg/assert"
	"github.com/andygeiss/machine-learning-classification/pkg/floats"
	"testing"
)

func TestVector_Max_Should_Handle_3_Values(t *testing.T) {
	assert.That(t, floats.Maximum([]float64{.5, 1.0, .75}), 1.0)
}

func TestVector_Max_Should_Handle_5_Values(t *testing.T) {
	assert.That(t, floats.Maximum([]float64{.5, 1.0, .75, 3.5, 0.4}), 3.5)
}

func TestVector_Max_Should_Handle_8_Values(t *testing.T) {
	assert.That(t, floats.Maximum([]float64{.5, 1.0, .75, 3.5, 0.4, 1.0, 2.0, 0.1}), 3.5)
}

func TestVector_Min_Should_Handle_3_Values(t *testing.T) {
	assert.That(t, floats.Minimum([]float64{1.5, 1.0, 2.75}), 1.0)
}

func TestVector_Min_Should_Handle_5_Values(t *testing.T) {
	assert.That(t, floats.Minimum([]float64{1.5, 1.0, 2.75, 3.5, 0.4}), 0.4)
}

func TestVector_Min_Should_Handle_8_Values(t *testing.T) {
	assert.That(t, floats.Minimum([]float64{1.5, 1.0, 2.75, 3.5, 0.4, 1.0, 2.0, 0.1}), 0.1)
}

func TestVector_Range_Should_Handle_3_Values(t *testing.T) {
	assert.That(t, floats.Range([]float64{1.5, 5.0, 6.5}), 5.0)
}
