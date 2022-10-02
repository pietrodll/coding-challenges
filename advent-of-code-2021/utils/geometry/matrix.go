package geometry

import "fmt"

type Matrix struct {
	values        [][]int
	height, width int
}

func MakeZeroMatrix(height, width int) *Matrix {
	values := make([][]int, height)

	for i := range values {
		values[i] = make([]int, width)
	}

	return &Matrix{values, height, width}
}

func MakeMatrix(data [][]int) *Matrix {
	height := len(data)
	width := len(data[0])

	values := make([][]int, height)

	for i := range values {
		values[i] = make([]int, width)
		copy(values[i], data[i])
	}

	return &Matrix{values, height, width}
}

func Identity(size int) *Matrix {
	result := MakeZeroMatrix(size, size)

	for i := 0; i < size; i++ {
		result.values[i][i] = 1
	}

	return result
}

func (m *Matrix) IsSquare() bool {
	return m.height == m.width
}

func (m *Matrix) Dimensions() (int, int) {
	return m.height, m.width
}

func (m *Matrix) Get(i, j int) int {
	return m.values[i][j]
}

func (m *Matrix) Add(other *Matrix) *Matrix {
	height, width := other.Dimensions()

	if height != m.height || width != m.width {
		panic(fmt.Errorf("cannot add matrix (%d, %d) and matrix (%d, %d)", m.height, m.width, height, width))
	}

	result := MakeZeroMatrix(m.height, m.width)

	for i, line := range result.values {
		for j := range line {
			result.values[i][j] = m.values[i][j] + other.values[i][j]
		}
	}

	return result
}

func (m *Matrix) Scale(x int) *Matrix {
	result := MakeZeroMatrix(m.height, m.width)

	for i, line := range result.values {
		for j := range line {
			result.values[i][j] = m.values[i][j] * x
		}
	}

	return result
}

func (m *Matrix) Sub(other *Matrix) *Matrix {
	if other.height != m.height || other.width != m.width {
		panic(fmt.Errorf("cannot substract matrix (%d, %d) and matrix (%d, %d)", m.height, m.width, other.height, other.width))
	}

	result := MakeZeroMatrix(m.height, m.width)

	for i, line := range result.values {
		for j := range line {
			result.values[i][j] = m.values[i][j] - other.values[i][j]
		}
	}

	return result
}

func (m *Matrix) Mult(other *Matrix) *Matrix {
	if other.height != m.width {
		panic(fmt.Errorf("cannot multiply matrix (%d, %d) and matrix (%d, %d)", m.height, m.width, other.height, other.width))
	}

	result := MakeZeroMatrix(m.height, other.width)

	for i, line := range result.values {
		for j := range line {
			for k, val := range m.values[i] {
				result.values[i][j] += val * other.values[k][j]
			}
		}
	}

	return result
}

func (m *Matrix) MultVec(v *Vector) *Vector {
	if v.size != m.width {
		panic(fmt.Errorf("cannot multiply matrix (%d, %d) and vector (%d)", m.height, m.width, v.size))
	}

	values := make([]int, m.height)

	for i := range values {
		for j, val := range v.values {
			values[i] += m.values[i][j] * val
		}
	}

	return MakeVector(values...)
}

func (m *Matrix) Transpose() *Matrix {
	result := MakeZeroMatrix(m.width, m.height)

	for i, line := range m.values {
		for j, val := range line {
			result.values[j][i] = val
		}
	}

	return result
}
