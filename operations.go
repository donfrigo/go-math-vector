package vector

import (
	"errors"
)

// Returns sum of two vectors
func (v1 *Vector) Add(v2 Vector) (Vector, error){
	err := v1.CompareDimensions(v2)
	if err != nil {
		return Vector{}, err
	}

	var sum []float64

	for i, val1 := range v1.Values {
		val2 := v2.Values[i]
		sum = append(sum, val1 + val2)
	}

	return New(sum), nil
}

// Returns difference of two vectors
func (v1 *Vector) Subtract(v2 Vector) (Vector, error){
	err := v1.CompareDimensions(v2)
	if err != nil {
		return Vector{}, err
	}

	var diff []float64

	for i, val1 := range v1.Values {
		val2 := v2.Values[i]
		diff = append(diff, val1 - val2)
	}

	return New(diff), nil
}

// Returns dot product of two vectors
func (v1 *Vector) DotProduct(v2 Vector) (float64, error){
	err := v1.CompareDimensions(v2)
	if err != nil {
		return 0.0, err
	}

	sum := 0.0

	for i, val1 := range v1.Values {
		val2 := v2.Values[i]
		sum += val1 * val2
	}

	return sum, nil
}

// Returns cross product of two, three-dimensional vectors
func (v1 *Vector) CrossProduct(v2 Vector) (Vector, error){
	if v1.Length() != 3 || v2.Length() != 3 {
		return Vector{}, errors.New("both vectors must be three-dimensional")
	}

	xProduct := make([]float64, 3)

	val1 := v1.Values
	val2 := v2.Values

	xProduct[0] = val1[1] * val2[2] - val1[2] * val2[1]
	xProduct[1] = val1[2] * val2[0] - val1[0] * val2[2]
	xProduct[2] = val1[0] * val2[1] - val1[1] * val2[0]

	return New(xProduct), nil
}

// Returns error, if dimensions don't match
// Would be possible to shrink the larger vector or add padding to the smaller vector, but we keep it basic for now
func (v1 *Vector) CompareDimensions(v2 Vector) error {
	if v1.Length() != v2.Length() {
		return errors.New("in order to carry out this operation, the dimensions of the two vectors should match")
	}
	return nil
}
