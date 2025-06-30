package distance

import (
	"math"
	"testing"
)

func TestCosineSimilarity(t *testing.T) {
	testCases := []struct {
		name      string
		a, b      []float64
		want      float64
		expectErr bool
	}{
		{
			name: "Perfectly similar vectors",
			a:    []float64{1, 2, 3},
			b:    []float64{1, 2, 3},
			want: 1,
		},
		{
			name: "Scaled similar vectors",
			a:    []float64{1, 2, 3},
			b:    []float64{2, 4, 6},
			want: 1,
		},
		{
			name: "Perfectly dissimilar vectors",
			a:    []float64{1, 2, 3},
			b:    []float64{-1, -2, -3},
			want: -1,
		},
		{
			name: "Orthogonal vectors",
			a:    []float64{1, 0},
			b:    []float64{0, 1},
			want: 0,
		},
		{
			name: "General case",
			a:    []float64{1, 2, 3, 4, 5},
			b:    []float64{5, 4, 3, 2, 1},
			// Dot product: 5+8+9+8+5 = 35
			// Norm A: sqrt(1+4+9+16+25) = sqrt(55)
			// Norm B: sqrt(25+16+9+4+1) = sqrt(55)
			// Similarity: 35 / (sqrt(55) * sqrt(55)) = 35 / 55 = 7 / 11
			want: 35.0 / 55.0,
		},
		{
			name:      "Mismatched lengths",
			a:         []float64{1, 2},
			b:         []float64{1, 2, 3},
			expectErr: true,
		},
		{
			name:      "Empty vectors",
			a:         []float64{},
			b:         []float64{},
			expectErr: true,
		},
		{
			name:      "Zero-magnitude vector",
			a:         []float64{0, 0, 0},
			b:         []float64{1, 2, 3},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CosineSimilarity(tc.a, tc.b)
			if (err != nil) != tc.expectErr {
				t.Fatalf("CosineSimilarity() error = %v, expectErr %v", err, tc.expectErr)
			}
			if !tc.expectErr {
				// Using a small tolerance for float comparison
				if math.Abs(got-tc.want) > 1e-9 {
					t.Errorf("CosineSimilarity() = %v, want %v", got, tc.want)
				}
			}
		})
	}
}
