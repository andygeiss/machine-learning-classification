package main

import (
	"flag"
	"fmt"
	"github.com/andygeiss/machine-learning-classification/internal/pipeline"
	"log"
	"path/filepath"
	"time"
)

func main() {

	k := flag.Int("k", 3, "Amount of nearest neighbors")
	x := flag.Float64("x", 0, "Vector X value")
	y := flag.Float64("y", 0, "Vector Y value")
	model := flag.String("model", filepath.Join("models", "iris_knn.pb"), "Model used for kNN")
	flag.Parse()

	start := time.Now()
	p := pipeline.New().
		PredictWithKNN(*x, *y, *k, *model)
	fmt.Printf("Prediction time: %v\n", time.Since(start))

	if err := p.Error(); err != nil {
		log.Fatal(err)
	}
}
