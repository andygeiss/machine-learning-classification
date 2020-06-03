package floats_test

import (
	"github.com/andygeiss/machine-learning-classification/pkg/assert"
	"github.com/andygeiss/machine-learning-classification/pkg/floats"
	goNum "gonum.org/v1/gonum/floats"
	"math/rand"
	"testing"
)

func TestSplitApplyCombine_Should_Handle_1000_Random_Values(t *testing.T) {
	values := randomValues(1000)
	pieces := &floats.Pieces{}
	r1 := floats.Maximum(values)
	r2 := pieces.SplitApplyCombine(values, floats.Maximum, floats.Maximum)
	assert.That(t, r1, r2)
}

func Benchmark_GoNum_Range(b *testing.B) {
	in := randomValues(1000000)
	for i := 0; i < b.N; i++ {
		min := goNum.Min(in)
		max := goNum.Max(in)
		_ = max - min
	}
}

func Benchmark_Floats_Range(b *testing.B) {
	in := randomValues(1000000)
	for i := 0; i < b.N; i++ {
		floats.Range(in)
	}
}

func Benchmark_Floats_SplitApplyCombine_Range(b *testing.B) {
	in := randomValues(1000000)
	pieces := &floats.Pieces{}
	for i := 0; i < b.N; i++ {
		max := pieces.SplitApplyCombine(in, floats.Maximum, floats.Maximum)
		min := pieces.SplitApplyCombine(in, floats.Minimum, floats.Minimum)
		_ = max - min
	}
}

func randomValues(size int) []float64 {
	values := make([]float64, size)
	for i := range values {
		values[i] = rand.Float64()
	}
	return values
}
