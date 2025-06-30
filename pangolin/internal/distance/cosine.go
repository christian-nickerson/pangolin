package distance

import (
	"fmt"
	"math"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
// It returns a score between -1 (perfectly dissimilar) and 1 (perfectly similar).
func CosineSimilarity(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vector lengths do not match: %d != %d", len(a), len(b))
	}

	if len(a) == 0 {
		return 0, fmt.Errorf("vectors cannot be empty")
	}

	dotProduct := 0.0
	normA := 0.0
	normB := 0.0

	for i := range len(a) {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	magnitude := math.Sqrt(normA) * math.Sqrt(normB)
	if magnitude == 0 {
		return 0, fmt.Errorf("cannot compute similarity with a zero-magnitude vector")
	}

	return dotProduct / magnitude, nil
}
