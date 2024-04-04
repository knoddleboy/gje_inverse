package gjeinverse

import (
	"time"
)

func InverseSerial(m *Matrix) (*Matrix, time.Duration, error) {
	n := m.Dim

	I := NewMatrix(n)
	I.FillIdentity()

	startTime := time.Now()

	for i := 0; i < n; i++ {
		if m.Data[i][i] == 0 {
			for j := i; j < n; j++ {
				if m.Data[j][i] != 0 {
					m.Swap(i, j)
					break
				}
				if j == n-1 {
					return nil, time.Since(startTime), ErrSingularMatrix
				}
			}
		}

		pivot := m.Data[i][i]

		for c := 0; c < n; c++ {
			m.Data[i][c] /= pivot
			I.Data[i][c] /= pivot
		}

		if i < n-1 {
			for r := i + 1; r < n; r++ {
				factor := m.Data[r][i]
				for c := 0; c < n; c++ {
					m.Data[r][c] -= factor * m.Data[i][c]
					I.Data[r][c] -= factor * I.Data[i][c]
				}
			}
		}
	}

	for i := n - 1; i >= 1; i-- {
		for r := i - 1; r >= 0; r-- {
			factor := m.Data[r][i]
			for c := 0; c < n; c++ {
				m.Data[r][c] -= factor * m.Data[i][c]
				I.Data[r][c] -= factor * I.Data[i][c]
			}
		}
	}

	return I, time.Since(startTime), nil
}
