package distance

import (
	"fmt"
	"math"

	"github.com/christian-nickerson/pangolin/pangolin/internal/models"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
func CosineSimilarity(a, b models.Vector) (float64, error) {
	if a.Length() != b.Length() {
		return 0, fmt.Errorf("vector lengths do not match: %d != %d", a.Length(), b.Length())
	}

	if a.Length() == 0 {
		return 0, fmt.Errorf("vectors cannot be empty")
	}

	dotProduct := 0.0
	normA := 0.0
	normB := 0.0

	for i := range a.Length() {
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
