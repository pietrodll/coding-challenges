package geometry

import "fmt"

type Vector struct {
	values []int
	size   int
}

func MakeVector(values ...int) *Vector {
	return &Vector{values, len(values)}
}

func MakeZeroVector(size int) *Vector {
	values := make([]int, size)
	return &Vector{values, size}
}

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) Get(i int) int {
	return v.values[i]
}

func (v *Vector) Eq(other *Vector) bool {
	if v.size != other.size {
		return false
	}

	for i, val := range v.values {
		if other.values[i] != val {
			return false
		}
	}

	return true
}

func (v *Vector) Add(other *Vector) *Vector {
	if v.size != other.size {
		panic(fmt.Errorf("cannot add vectors of different size (%d, %d)", v.size, other.size))
	}

	result := MakeZeroVector(v.size)

	for i := range result.values {
		result.values[i] = v.values[i] + other.values[i]
	}

	return result
}

func (v *Vector) Sub(other *Vector) *Vector {
	if v.size != other.size {
		panic(fmt.Errorf("cannot substract vectors of different sizes (%d, %d)", v.size, other.size))
	}

	result := MakeZeroVector(v.size)

	for i := range result.values {
		result.values[i] = v.values[i] - other.values[i]
	}

	return result
}

func (v *Vector) Dot(other *Vector) int {
	if v.size != other.size {
		panic(fmt.Errorf("cannot calculate dot product between vectors of different sizes (%d, %d)", v.size, other.size))
	}

	result := 0

	for i := range v.values {
		result += v.values[i] * other.values[i]
	}

	return result
}
