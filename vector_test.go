package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	values := []float64{1, 1, 1}

	v1 := New(values)
	assert.EqualValues(v1.Values, values)
}

func TestVector_Length(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1})
	expected := 3

	assert.EqualValues(v1.Length(), expected)
}

func TestVector_Magnitude(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 2, 2})
	expected := 3

	assert.EqualValues(v1.Magnitude(), expected)
}

func TestVector_ScalarMultiplication(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1})
	v1 = v1.ScalarMultiplication(3)
	expected := []float64{3, 3, 3}

	assert.EqualValues(expected, v1.Values)
}
