package floats_test

import (
	"github.com/andygeiss/machine-learning-classification/pkg/assert"
	"github.com/andygeiss/machine-learning-classification/pkg/floats"
	"testing"
)

func TestReplace_Should_Handle_One_Value(t *testing.T) {
	in := []float64{1}
	floats.Replace(in, 1, 2)
	assert.That(t, in[0], 2)
}

func TestReplace_Should_Handle_Two_Values(t *testing.T) {
	in := []float64{1, 1}
	floats.Replace(in, 1, 2)
	assert.That(t, in[0], 2)
	assert.That(t, in[1], 2)
}

func TestReplace_Should_Handle_Three_Values(t *testing.T) {
	in := []float64{1, 1, 1}
	floats.Replace(in, 1, 2)
	assert.That(t, in[0], 2)
	assert.That(t, in[1], 2)
	assert.That(t, in[2], 2)
}

func TestReplace_Should_Handle_Four_Values(t *testing.T) {
	in := []float64{1, 3, 1, 1}
	floats.Replace(in, 1, 2)
	assert.That(t, in[0], 2)
	assert.That(t, in[1], 3)
	assert.That(t, in[2], 2)
	assert.That(t, in[3], 2)
}
