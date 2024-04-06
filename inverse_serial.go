package gjeinverse

import (
	"time"
)

func InverseSerial(a *Matrix) (*Matrix, time.Duration) {
	n := a.Dim

	I := NewMatrix(n)
	I.FillIdentity()

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
			for r := i + 1; r < n; r++ {
				factor := a.Data[r][i]
				subtractRows(a.Data[r], a.Data[i], factor)
				subtractRows(I.Data[r], I.Data[i], factor)
			}
		}
	}

	for i := n - 1; i >= 1; i-- {
		for r := i - 1; r >= 0; r-- {
			factor := a.Data[r][i]
			subtractRows(a.Data[r], a.Data[i], factor)
			subtractRows(I.Data[r], I.Data[i], factor)
		}
	}

	return I, time.Since(startTime)
}
