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

// NewMatrix creates a new square matrix of dimension 'dim'.
// It panics if 'dim' is less than or equal to zero.
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

// Randomize fills the matrix with random float64 values.
func (a *Matrix) Randomize() {
	for i := 0; i < a.Dim; i++ {
		for j := 0; j < a.Dim; j++ {
			a.Data[i][j] = rand.Float64()
		}
	}
}

// FillIdentity fills the matrix as an identity matrix.
func (a *Matrix) FillIdentity() {
	for i := 0; i < a.Dim; i++ {
		for j := 0; j < a.Dim; j++ {
			if i == j {
				a.Data[i][j] = 1
			} else {
				a.Data[i][j] = 0
			}
		}
	}
}

// Copy returns a deep copy of the matrix.
func (a *Matrix) Copy() *Matrix {
	n := a.Dim
	copied := NewMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			copied.Data[i][j] = a.Data[i][j]
		}
	}
	return copied
}

// Swap swaps two rows of the matrix.
func (a *Matrix) Swap(i, j int) {
	a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
}

// IsIdentity checks if the matrix is an identity matrix.
func (a *Matrix) IsIdentity() bool {
	for i := 0; i < a.Dim; i++ {
		for j := 0; j < a.Dim; j++ {
			if i == j {
				if a.Data[i][j] != 1 {
					return false
				}
			} else {
				if a.Data[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}

// Equals checks if two matrices are equal within a certain tolerance.
func (a *Matrix) Equals(o *Matrix) bool {
	n := a.Dim
	tol := 1e-8
	if n != o.Dim {
		return false
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			diff := math.Abs(a.Data[r][c] - o.Data[r][c])
			if diff > tol {
				return false
			}
		}
	}
	return true
}

// Print prints the matrix in a tabular format.
func (a *Matrix) Print() {
	for i := 0; i < a.Dim; i++ {
		for j := 0; j < a.Dim; j++ {
			fmt.Printf("%11f", a.Data[i][j])
		}
		fmt.Println()
	}
}

// Printf prints the matrix in a formatted style.
// The output can be used in https://www.wolframalpha.com inverse matrix calculator.
func (a *Matrix) Printf() {
	fmt.Print("{")
	for i := 0; i < a.Dim; i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("{")
		for j := 0; j < a.Dim; j++ {
			fmt.Printf("%f", a.Data[i][j])
			if j < a.Dim-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("}")
	}
	fmt.Println("}")
}
