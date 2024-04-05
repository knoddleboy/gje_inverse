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
		err := assertNonZeroPivot(m, I, i)
		if err != nil {
			panic(err)
		}

		pivot := m.Data[i][i]
		normalizeRow(m.Data[i], pivot)
		normalizeRow(I.Data[i], pivot)

		if i < n-1 {
			for r := i + 1; r < n; r++ {
				factor := m.Data[r][i]
				subtractRows(m.Data[r], m.Data[i], factor)
			}
		}
	}

	for i := n - 1; i >= 1; i-- {
		for r := i - 1; r >= 0; r-- {
			factor := m.Data[r][i]
			subtractRows(m.Data[r], m.Data[i], factor)
		}
	}

	return I, time.Since(startTime), nil
}
