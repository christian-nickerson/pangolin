package models

type Vector []float64

// Length returns the length of the vector
func (x Vector) Length() int {
	return len(x)
}
