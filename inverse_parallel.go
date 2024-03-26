package gjeinverse

import (
	"time"
)

func InverseParallel(m *Matrix, threads int) (*Matrix, time.Duration, error) {
	n := m.Dim

	I := NewMatrix(n)
	I.FillIdentity()

	ch := make(chan int, threads)

	startTime := time.Now()

	for i := 0; i < n; i++ {
		if m.Data[i][i] == 0 {
			for j := i + 1; j < n; j++ {
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

		for gr := 0; gr < threads; gr++ {
			go func(startRow, endRow int) {
				for r := startRow; r < endRow; r++ {
					if r != i {
						factor := m.Data[r][i]
						for c := 0; c < n; c++ {
							m.Data[r][c] -= factor * m.Data[i][c]
							I.Data[r][c] -= factor * I.Data[i][c]
						}
					}
				}
				ch <- 1
			}(gr*n/threads, (gr+1)*n/threads)
		}

		for gr := 0; gr < threads; gr++ {
			<-ch
		}
	}

	return I, time.Since(startTime), nil
}
