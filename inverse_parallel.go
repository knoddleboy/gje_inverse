package gjeinverse

import (
	"sync"
	"time"
)

func InverseParallel(a *Matrix, threads int) (*Matrix, time.Duration, error) {
	n := a.Dim

	I := NewMatrix(n)
	I.FillIdentity()

	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < n; i++ {
		err := assertNonZeroPivot(a, I, i)
		if err != nil {
			panic(err)
		}

		pivot := a.Data[i][i]
		normalizeRow(a.Data[i], pivot)
		normalizeRow(I.Data[i], pivot)

		if i < n-1 {
			wg.Add(n - i - 1)
			for r := i + 1; r < n; r++ {
				go func(row int) {
					defer wg.Done()
					factor := a.Data[row][i]
					subtractRows(a.Data[row], a.Data[i], factor)
					subtractRows(I.Data[row], I.Data[i], factor)
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
				factor := a.Data[r][col]
				subtractRows(a.Data[r], a.Data[col], factor)
				subtractRows(I.Data[r], I.Data[col], factor)
			}(i)
		}
		wg.Wait()
	}

	return I, time.Since(startTime), nil
}
