package floats_test

import (
	"github.com/andygeiss/machine-learning-classification/pkg/assert"
	"github.com/andygeiss/machine-learning-classification/pkg/floats"
	"testing"
)

func TestMinMaxScale(t *testing.T) {
	in := []float64{115, 140, 175}
	floats.MinMaxScale(in)
	assert.That(t, in[0], 0)
	assert.That(t, in[1], 0.4166666666666667)
	assert.That(t, in[2], 1)
}

func TestStandardScale(t *testing.T) {
	in := []float64{115, 140, 175}
	floats.StandardScale(in)
	assert.That(t, in[0], -1.151385284513614)
	assert.That(t, in[1], -0.13545709229571964)
	assert.That(t, in[2], 1.2868423768093327)
}
