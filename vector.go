package vector

import "math"

type Vector struct {
	Values []float64
}

// Returns object of type Vector
func New (n []float64) Vector {
	vector := Vector{
		Values: n,
	}

	return vector
}

// Returns Length of Vector v1
func (v1 *Vector) Length() int{
	return len(v1.Values)
}

// Returns Magnitude of Vector v1
func (v1 *Vector) Magnitude() float64{
	magnitude := 0.0
	for _, val := range v1.Values {
		magnitude += math.Pow(val, 2)
	}
	return math.Sqrt(magnitude)
}

// Returns Scalar multiplication of Vector v1
func (v1 *Vector) ScalarMultiplication(n float64) Vector{
	res := make([]float64, v1.Length())
	for i, val := range v1.Values {
		res[i] = val * n
	}
	return New(res)
}
