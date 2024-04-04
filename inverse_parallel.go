package gjeinverse

import (
	"sync"
	"time"
)

func InverseParallel(m *Matrix, threads int) (*Matrix, time.Duration, error) {
	n := m.Dim

	I := NewMatrix(n)
	I.FillIdentity()

	var wg sync.WaitGroup

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
			wg.Add(n - i - 1)
			for r := i + 1; r < n; r++ {
				go func(row int) {
					defer wg.Done()
					factor := m.Data[row][i]
					for c := 0; c < n; c++ {
						m.Data[row][c] -= factor * m.Data[i][c]
						I.Data[row][c] -= factor * I.Data[i][c]
					}
				}(r)
			}
			wg.Wait()
		}
	}

	wg.Add(n - 1)
	for i := n - 1; i >= 1; i-- {
		go func(i int) {
			defer wg.Done()
			for r := i - 1; r >= 0; r-- {
				factor := m.Data[r][i]
				for c := 0; c < n; c++ {
					m.Data[r][c] -= factor * m.Data[i][c]
					I.Data[r][c] -= factor * I.Data[i][c]
				}
			}
		}(i)
	}
	wg.Wait()

	return I, time.Since(startTime), nil
}
