package distance

import (
	"math"
	"testing"

	"github.com/christian-nickerson/pangolin/pangolin/internal/models"
)

func TestEuclideanDistance(t *testing.T) {
	testCases := []struct {
		name      string
		a, b      models.Vector
		want      float64
		expectErr bool
	}{
		{
			name: "Simple case",
			a:    []float64{1, 2, 3},
			b:    []float64{4, 5, 6},
			want: math.Sqrt(27),
		},
		{
			name: "Same vector",
			a:    []float64{1, 2, 3},
			b:    []float64{1, 2, 3},
			want: 0,
		},
		{
			name: "Negative values",
			a:    []float64{-1, -2, -3},
			b:    []float64{-4, -5, -6},
			want: math.Sqrt(27),
		},
		{
			name:      "Mismatched lengths",
			a:         []float64{1, 2, 3},
			b:         []float64{4, 5},
			expectErr: true,
		},
		{
			name:      "Empty vectors",
			a:         []float64{},
			b:         []float64{},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := EuclideanDistance(tc.a, tc.b)
			if (err != nil) != tc.expectErr {
				t.Fatalf("EuclideanDistance() error = %v, expectErr %v", err, tc.expectErr)
			}
			if !tc.expectErr && got != tc.want {
				t.Errorf("EuclideanDistance() = %v, want %v", got, tc.want)
			}
		})
	}
}
