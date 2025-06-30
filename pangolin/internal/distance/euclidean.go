package distance

import (
	"fmt"
	"math"

	"github.com/christian-nickerson/pangolin/pangolin/internal/models"
)

// EuclideanDistance calculates the Euclidean distance between two vectors.
func EuclideanDistance(a, b models.Vector) (float64, error) {
	if a.Length() != b.Length() {
		return 0, fmt.Errorf("vector lengths do not match: %d != %d", a.Length(), b.Length())
	}

	if a.Length() == 0 {
		return 0, fmt.Errorf("vectors cannot be empty")
	}

	var sumOfSquares float64
	for i := range a {
		diff := a[i] - b[i]
		sumOfSquares += diff * diff
	}

	return math.Sqrt(sumOfSquares), nil
}
