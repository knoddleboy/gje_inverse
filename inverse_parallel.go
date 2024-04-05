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
		err := assertNonZeroPivot(m, I, i)
		if err != nil {
			panic(err)
		}

		pivot := m.Data[i][i]
		normalizeRow(m.Data[i], pivot)
		normalizeRow(I.Data[i], pivot)

		if i < n-1 {
			wg.Add(n - i - 1)
			for r := i + 1; r < n; r++ {
				go func(row int) {
					defer wg.Done()
					factor := m.Data[row][i]
					subtractRows(m.Data[row], m.Data[i], factor)
				}(r)
			}
			wg.Wait()
		}
	}

	for i := n - 1; i >= 1; i-- {
		wg.Add(i)
		for r := i - 1; r >= 0; r-- {
			go func(col int) {
				defer wg.Done()
				factor := m.Data[r][col]
				subtractRows(m.Data[r], m.Data[col], factor)
			}(i)
		}
		wg.Wait()
	}

	return I, time.Since(startTime), nil
}
