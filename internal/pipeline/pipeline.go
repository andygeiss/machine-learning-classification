package pipeline

import (
	"encoding/csv"
	"fmt"
	"github.com/andygeiss/machine-learning-classification/internal/api"
	"github.com/andygeiss/machine-learning-classification/pkg/floats"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type pipeline struct {
	err error
}

func (p *pipeline) Error() error {
	return p.err
}

func (p *pipeline) EvaluateWithKnn(k int, inFilename string) *pipeline {
	if p.err != nil {
		return p
	}

	pb, err := ioutil.ReadFile(inFilename)
	if err != nil {
		p.err = err
		return p
	}

	var data api.ProcessedData
	if err := proto.Unmarshal(pb, &data); err != nil {
		p.err = err
		return p
	}

	fmt.Printf("K-Nearest Neighbour with k = %d :\n", k)

	correct := 0
	count := 0
	for i := range data.Species {
		given := []float64{data.PetalLength[i], data.PetalWidth[i]}
		wanted := data.Species[i]
		predicted := floats.Knn(given, data.PetalLength, data.PetalWidth, data.Species, k, floats.ManhattanDistance)
		if wanted == predicted {
			correct++
		}
		count++
	}
	fmt.Printf("   Petal Length/Width Accuracy: %2.2f\n", float64(correct*100/count))

	correct = 0
	count = 0
	for i := range data.Species {
		given := []float64{data.SepalLength[i], data.SepalWidth[i]}
		wanted := data.Species[i]
		predicted := floats.Knn(given, data.SepalLength, data.SepalWidth, data.Species, k, floats.ManhattanDistance)
		if wanted == predicted {
			correct++
		}
		count++
	}
	fmt.Printf("   Sepal Length/Width Accuracy: %2.2f\n", float64(correct*100/count))

	return p
}

func (p *pipeline) GatherData(sourceURL, targetFilename string) *pipeline {
	if p.err != nil {
		return p
	}

	res, err := http.Get(sourceURL)
	if err != nil {
		p.err = err
		return p
	}
	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		p.err = err
		return p
	}

	p.err = ioutil.WriteFile(targetFilename, raw, 0644)

	return p
}

func (p *pipeline) MeasureStatistics(processedFilename string, ) *pipeline {

	if p.err != nil {
		return p
	}

	pb, err := ioutil.ReadFile(processedFilename)
	if err != nil {
		p.err = err
		return p
	}

	var data api.ProcessedData
	if err := proto.Unmarshal(pb, &data); err != nil {
		p.err = err
		return p
	}

	fmt.Printf("Statistics:\n")
	fmt.Printf("   %-16s %-8s %-8s %-8s %-8s %-8s %-8s %-8s %-8s\n", "Column", "Mean", "Median", "Mode", "Minimum", "Maximum", "Range", "Variance", "Std Dev")
	fmt.Printf("   %-16s %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f\n", "Petal length", floats.Mean(data.PetalLength), floats.Median(data.PetalLength), floats.Mode(data.PetalLength), floats.Minimum(data.PetalLength), floats.Maximum(data.PetalLength), floats.Range(data.PetalLength), floats.Variance(data.PetalLength), floats.StandardDeviation(data.PetalLength))
	fmt.Printf("   %-16s %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f\n", "Petal width", floats.Mean(data.PetalWidth), floats.Median(data.PetalWidth), floats.Mode(data.PetalWidth), floats.Minimum(data.PetalWidth), floats.Maximum(data.PetalWidth), floats.Range(data.PetalWidth), floats.Variance(data.PetalWidth), floats.StandardDeviation(data.PetalWidth))
	fmt.Printf("   %-16s %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f\n", "Sepal length", floats.Mean(data.SepalLength), floats.Median(data.SepalLength), floats.Mode(data.SepalLength), floats.Minimum(data.SepalLength), floats.Maximum(data.SepalLength), floats.Range(data.SepalLength), floats.Variance(data.SepalLength), floats.StandardDeviation(data.SepalLength))
	fmt.Printf("   %-16s %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f %-8.2f\n", "Sepal width", floats.Mean(data.SepalWidth), floats.Median(data.SepalWidth), floats.Mode(data.SepalWidth), floats.Minimum(data.SepalLength), floats.Maximum(data.SepalLength), floats.Range(data.SepalLength), floats.Variance(data.SepalLength), floats.StandardDeviation(data.SepalLength))

	return p
}

func (p *pipeline) NormalizeFeatures(inFilename, outFilename string) *pipeline {
	if p.err != nil {
		return p
	}

	pb, err := ioutil.ReadFile(inFilename)
	if err != nil {
		p.err = err
		return p
	}

	var data api.ProcessedData
	if err := proto.Unmarshal(pb, &data); err != nil {
		p.err = err
		return p
	}

	floats.MinMaxScale(data.PetalLength)
	floats.MinMaxScale(data.PetalWidth)
	floats.MinMaxScale(data.SepalLength)
	floats.MinMaxScale(data.SepalWidth)

	out, err := proto.Marshal(&data)
	if err != nil {
		p.err = err
		return p
	}

	if err := ioutil.WriteFile(outFilename, out, 0644); err != nil {
		p.err = err
		return p
	}

	return p
}

func (p *pipeline) OrganizeData(interimFilename, processedFilename string) *pipeline {
	if p.err != nil {
		return p
	}

	pb, err := ioutil.ReadFile(interimFilename)
	if err != nil {
		p.err = err
		return p
	}

	var interimData api.InterimData
	if err := proto.Unmarshal(pb, &interimData); err != nil {
		p.err = err
		return p
	}

	n := len(interimData.Records)
	processedData := new(api.ProcessedData)
	processedData.PetalLength = make([]float64, n)
	processedData.PetalWidth = make([]float64, n)
	processedData.SepalLength = make([]float64, n)
	processedData.SepalWidth = make([]float64, n)
	processedData.Species = make([]float64, n)
	processedData.Size = float64(n)
	processedData.Timestamp = int64(time.Now().Second())

	// Map the species string to a fixed float64 number.
	speciesFloat64 := map[string]float64{
		"Iris-setosa":     1,
		"Iris-versicolor": 2,
		"Iris-virginica":  3,
	}

	for i, record := range interimData.Records {
		processedData.PetalLength[i] = record.PetalLength
		processedData.PetalWidth[i] = record.PetalWidth
		processedData.SepalLength[i] = record.SepalLength
		processedData.SepalWidth[i] = record.SepalWidth
		processedData.Species[i] = speciesFloat64[record.Species]
	}

	out, err := proto.Marshal(processedData)
	if err != nil {
		p.err = err
		return p
	}

	if err := ioutil.WriteFile(processedFilename, out, 0644); err != nil {
		p.err = err
		return p
	}

	return p
}

func (p *pipeline) PredictWithKNN(x, y float64, k int, inFilename string) *pipeline {

	if p.err != nil {
		return p
	}

	pb, err := ioutil.ReadFile(inFilename)
	if err != nil {
		p.err = err
		return p
	}

	var data api.ProcessedData
	if err := proto.Unmarshal(pb, &data); err != nil {
		p.err = err
		return p
	}

	// Map the species string to a fixed float64 number.
	species := map[float64]string{
		1: "Iris-setosa",
		2: "Iris-versicolor",
		3: "Iris-virginica",
	}

	given := []float64{x, y}

	floats.MinMaxScale(given)

	predicted := floats.Knn(given, data.SepalLength, data.SepalWidth, data.Species, k, floats.ManhattanDistance)
	fmt.Printf("K-Nearest Neighbour - K: %d, Given: %v, Predicted: %s\n", k, given, species[predicted])

	return p
}

func (p *pipeline) ReplaceZerosToMean(inFilename, outFilename string) *pipeline {
	if p.err != nil {
		return p
	}

	pb, err := ioutil.ReadFile(inFilename)
	if err != nil {
		p.err = err
		return p
	}

	var data api.ProcessedData
	if err := proto.Unmarshal(pb, &data); err != nil {
		p.err = err
		return p
	}

	floats.Replace(data.PetalLength, 0, floats.Mean(data.PetalLength))
	floats.Replace(data.PetalWidth, 0, floats.Mean(data.PetalWidth))
	floats.Replace(data.SepalLength, 0, floats.Mean(data.SepalLength))
	floats.Replace(data.SepalWidth, 0, floats.Mean(data.SepalWidth))

	out, err := proto.Marshal(&data)
	if err != nil {
		p.err = err
		return p
	}

	if err := ioutil.WriteFile(outFilename, out, 0644); err != nil {
		p.err = err
		return p
	}

	return p
}

func (p *pipeline) TransformData(csvFilename, protoFilename string) *pipeline {
	if p.err != nil {
		return p
	}

	file, err := os.Open(csvFilename)
	if err != nil {
		p.err = err
		return p
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 5

	records, err := reader.ReadAll()
	if err != nil {
		p.err = err
		return p
	}

	input := new(api.InterimData)
	input.Records = make([]*api.InterimData_Record, 0)

	for _, rec := range records {

		dr := new(api.InterimData_Record)

		petalLength, err := strconv.ParseFloat(rec[0], 64)
		if err != nil {
			p.err = err
			return p
		}

		petalWidth, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			p.err = err
			return p
		}

		sepalLength, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			p.err = err
			return p
		}

		sepalWidth, err := strconv.ParseFloat(rec[3], 64)
		if err != nil {
			p.err = err
			return p
		}

		dr.PetalLength = petalLength
		dr.PetalWidth = petalWidth
		dr.SepalLength = sepalLength
		dr.SepalWidth = sepalWidth
		dr.Species = rec[4]

		input.Records = append(input.Records, dr)
	}

	out, err := proto.Marshal(input)
	if err != nil {
		p.err = err
		return p
	}

	p.err = ioutil.WriteFile(protoFilename, out, 0644)

	return p
}

func New() *pipeline {
	return &pipeline{}
}
