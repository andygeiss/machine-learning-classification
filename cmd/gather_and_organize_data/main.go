package main

import (
	"github.com/andygeiss/machine-learning-classification/internal/pipeline"
	"log"
	"path/filepath"
)

func main() {

	p := pipeline.New()

	sourceUrl := "https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data"
	externalData := filepath.Join("data", "external", "iris.csv")
	interimData := filepath.Join("data", "interim", "iris_interim.pb")
	processedData := filepath.Join("data", "processed", "iris_processed.pb")

	p.
		GatherData(sourceUrl, externalData).
		TransformData(externalData, interimData).
		OrganizeData(interimData, processedData).
		MeasureStatistics(processedData)

	if err := p.Error(); err != nil {
		log.Fatal(err)
	}
}
