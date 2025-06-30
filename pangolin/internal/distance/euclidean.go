package distance

import (
	"fmt"
	"math"
)

// EuclideanDistance calculates the Euclidean distance between two vectors.
func EuclideanDistance(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vector lengths do not match: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return 0, fmt.Errorf("vectors cannot be empty")
	}

	var sumOfSquares float64
	for i := range a {
		diff := a[i] - b[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares), nil
}
