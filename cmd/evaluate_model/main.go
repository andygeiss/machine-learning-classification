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

	k := flag.Int("k", 3, "Amount of Nearest Neighbors")
	flag.Parse()

	p := pipeline.New().
		ReplaceZerosToMean(filepath.Join("data", "processed", "iris_processed.pb"), filepath.Join("models", "iris_knn.pb")).
		NormalizeFeatures(filepath.Join("models", "iris_knn.pb"), filepath.Join("models", "iris_knn.pb")).
		MeasureStatistics(filepath.Join("models", "iris_knn.pb"))

	start := time.Now()
	p.EvaluateWithKnn(*k, filepath.Join("models", "iris_knn.pb"))
	fmt.Printf("Evaluation time: %v\n", time.Since(start))

	if err := p.Error(); err != nil {
		log.Fatal(err)
	}
}
