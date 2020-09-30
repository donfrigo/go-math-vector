package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector_Add(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1})
	v2 := New([]float64{-1, -1, -1})

	expected := []float64{0, 0, 0}
	sum, err := v1.Add(v2)

	assert.Nil(err)
	assert.EqualValues(sum.Values, expected)
}
func TestVector_AddDifferentDimensions(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1, 1})
	v2 := New([]float64{-1, -1, -1})

	expected := []float64{0, 0, 0}
	sum, err := v1.Add(v2)

	assert.NotNil(err)
	assert.NotEqual(expected, sum)
}

func TestVector_Subtract(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1})
	v2 := New([]float64{-1, -1, -1})

	expected := []float64{2, 2, 2}
	sum, err := v1.Subtract(v2)

	assert.Nil(err)
	assert.EqualValues(sum.Values, expected)
}
func TestVector_SubtractDifferentDimensions(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1, 1})
	v2 := New([]float64{-1, -1, -1})

	expected := []float64{0, 0, 0}
	sum, err := v1.Subtract(v2)

	assert.NotNil(err)
	assert.NotEqual(expected, sum)
}

func TestVector_CompareDimensions(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1})
	v2 := New([]float64{-1, -1, -1})

	err := v1.CompareDimensions(v2)

	assert.Nil(err)
}
func TestVector_CompareDimensionsDifferent(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 1, 1, 1})
	v2 := New([]float64{-1, -1, -1})

	err := v1.CompareDimensions(v2)

	assert.NotNil(err)
}

func TestVector_CrossProduct(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{3, -3, 1})
	v2 := New([]float64{-12, 12, -4})

	product, err := v1.CrossProduct(v2)
	expected := []float64{0, 0, 0}

	assert.Nil(err)
	assert.EqualValues(product.Values, expected)
}
func TestVector_CrossProduct4D(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{3, -3, 1, 3})
	v2 := New([]float64{-12, 12, -4, 4})

	product, err := v1.CrossProduct(v2)
	expected := []float64{0, 0, 0}

	assert.NotNil(err)
	assert.NotEqual(product.Values, expected)
}

func TestVector_DotProduct(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 2, 3})
	v2 := New([]float64{4, -5, 2})

	product, err := v1.DotProduct(v2)
	expected := 0.0

	assert.Nil(err)
	assert.Equal(product, expected)
}
func TestVector_DotProductDifferentDimensions(t *testing.T) {
	assert := assert.New(t)

	v1 := New([]float64{1, 2, 3, 3})
	v2 := New([]float64{4, -5, 2})

	_, err := v1.DotProduct(v2)

	assert.NotNil(err)
}




