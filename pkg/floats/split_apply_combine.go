package floats

import (
	"runtime"
	"sync"
)

// Pieces offers concurrent access to each of goroutine during apply state.
type Pieces struct {
	results []float64
	mutex   sync.Mutex
}

// SplitApplyCombine break up a slice of float64 data into smaller pieces (split),
// operate on each piece independently by a goroutine (apply)
// and put all the pieces back together (combine).
func (p *Pieces) SplitApplyCombine(in []float64, apply func(in []float64) (out float64), combine func(in []float64) (out float64)) (out float64) {
	// For CPU-bound workload its enough to use exactly one separate goroutines per CPU-thread.
	threads := maxThreads()
	// Calculate the size per piece which will be a multiple of CPU-threads.
	p.results = make([]float64, 0)
	sizePerPiece := len(in) / threads
	// Check if size is smaller than the amount of CPU-threads.
	// Only use the goroutine overhead on larger slices.
	if sizePerPiece > 0 {
		// We will wait for all CPU-threads.
		var wg sync.WaitGroup
		wg.Add(threads)
		// Now do the work and create a goroutine for each piece (Split).
		for i := 0; i < threads; i++ {
			// Make i visible for the following goroutine.
			i := i
			go func() {
				defer wg.Done()
				// Use i to select the start and end of the piece
				start := i * sizePerPiece
				end := (i + 1) * sizePerPiece
				// Work on the piece (Apply).
				piece := apply(in[start:end])
				// Add the result.
				p.mutex.Lock()
				defer p.mutex.Unlock()
				p.results = append(p.results, piece)
			}()
		}
		wg.Wait()
	}
	// Calculate the remaining tail of values.
	tail := len(in) % threads
	// Work on the tail if it's there.
	if tail > 0 {
		start := len(p.results) * sizePerPiece
		end := len(p.results) + 1*tail
		p.mutex.Lock()
		defer p.mutex.Unlock()
		p.results = append(p.results, apply(in[start:end]))
	}
	// Finally combine the results from each piece.
	return combine(p.results)
}

func maxThreads() int {
	maxCpus := runtime.GOMAXPROCS(0)
	numCpus := runtime.NumCPU()
	if maxCpus < numCpus {
		return maxCpus
	}
	return numCpus
}
