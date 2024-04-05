package gjeinverse

import (
	"fmt"
	"math"
	"math/rand"
)

type Matrix struct {
	Data [][]float64
	Dim  int
}

func NewMatrix(dim int) *Matrix {
	if dim <= 0 {
		panic(ErrNegativeDimension)
	}
	data := make([][]float64, dim)
	for i := 0; i < dim; i++ {
		data[i] = make([]float64, dim)
	}
	return &Matrix{
		Dim:  dim,
		Data: data,
	}
}

func (m *Matrix) Randomize() {
	for i := 0; i < m.Dim; i++ {
		for j := 0; j < m.Dim; j++ {
			m.Data[i][j] = rand.Float64()
		}
	}
}

func (m *Matrix) FillIdentity() {
	for i := 0; i < m.Dim; i++ {
		for j := 0; j < m.Dim; j++ {
			if i == j {
				m.Data[i][j] = 1
			} else {
				m.Data[i][j] = 0
			}
		}
	}
}

func (m *Matrix) Copy() *Matrix {
	n := m.Dim
	copied := NewMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			copied.Data[i][j] = m.Data[i][j]
		}
	}
	return copied
}

func (m *Matrix) Swap(i, j int) {
	m.Data[i], m.Data[j] = m.Data[j], m.Data[i]
}

func (m *Matrix) IsIdentity() bool {
	for i := 0; i < m.Dim; i++ {
		for j := 0; j < m.Dim; j++ {
			if i == j {
				if m.Data[i][j] != 1 {
					return false
				}
			} else {
				if m.Data[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func (m *Matrix) Equals(o *Matrix) bool {
	n := m.Dim
	if n != o.Dim {
		return false
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			diff := math.Abs(m.Data[r][c] - o.Data[r][c])
			if diff > 1e-6 {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Print() {
	for i := 0; i < m.Dim; i++ {
		for j := 0; j < m.Dim; j++ {
			fmt.Printf("%11f", m.Data[i][j])
		}
		fmt.Println()
	}
}

func (m *Matrix) Printf() {
	fmt.Print("{")
	for i := 0; i < m.Dim; i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("{")
		for j := 0; j < m.Dim; j++ {
			fmt.Printf("%f", m.Data[i][j])
			if j < m.Dim-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("}")
	}
	fmt.Println("}")
}
