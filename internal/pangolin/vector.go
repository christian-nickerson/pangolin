package pangolin

import (
	"fmt"
	"math"
)

// Vector represents a point in a high-dimensional space.
type Vector []float64

// VectorStore holds a collection of vectors.
type VectorStore struct {
	vectors map[string]Vector
}

// NewVectorStore creates a new VectorStore.
func NewVectorStore() *VectorStore {
	return &VectorStore{
		vectors: make(map[string]Vector),
	}
}

// AddVector adds a vector to the store with a given ID.
func (vs *VectorStore) AddVector(id string, vector Vector) {
	vs.vectors[id] = vector
}

// Search finds the top N most similar vectors to a query vector.
func (vs *VectorStore) Search(query Vector, topN int) ([]string, error) {
	if len(vs.vectors) == 0 {
		return nil, fmt.Errorf("vector store is empty")
	}

	scores := make(map[string]float64)
	for id, vector := range vs.vectors {
		score, err := cosineSimilarity(query, vector)
		if err != nil {
			return nil, err
		}
		scores[id] = score
	}

	// Simple (and inefficient) sorting to find top N.
	// For a real application, a more efficient method like a heap would be used.
	var resultIDs []string
	for i := 0; i < topN && i < len(scores); i++ {
		var maxID string
		var maxScore float64 = -1

		for id, score := range scores {
			if score > maxScore {
				maxScore = score
				maxID = id
			}
		}

		if maxID != "" {
			resultIDs = append(resultIDs, maxID)
			delete(scores, maxID)
		}
	}

	return resultIDs, nil
}

// cosineSimilarity calculates the cosine similarity between two vectors.
func cosineSimilarity(a, b Vector) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vectors must have the same dimension")
	}

	var dotProduct, normA, normB float64
	for i := 0; i < len(a); i++ {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0, fmt.Errorf("cannot compute similarity with a zero vector")
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB)), nil
}
