package gjeinverse

// assertNonZeroPivot ensures that the pivot element in the given matrix m at the specified row is non-zero.
// If the pivot is zero, it searches for a non-zero element in the same column below the current row,
// swaps rows if found, and updates the corresponding rows in the identity matrix I.
// Returns an error ErrSingularMatrix if no non-zero element is found below the current row.
func assertNonZeroPivot(a, I *Matrix, r int) error {
	n := a.Dim
	if a.Data[r][r] == 0 {
		for j := r; j < n; j++ {
			if a.Data[j][r] != 0 {
				a.Swap(r, j)
				I.Swap(r, j)
				break
			}
			if j == n-1 {
				return ErrSingularMatrix
			}
		}
	}
	return nil
}

// normalizeRow divides each element of the given row slice by the provided pivot value.
func normalizeRow(r []float64, pivot float64) {
	n := len(r)
	for c := 0; c < n; c++ {
		r[c] /= pivot
	}
}

// subtractRows subtracts row-2 (r2) from row-1 (r1) scaled by the provided factor value.
// For example, if r1 = (1 2 3) and r2 = (4 5 6) with factor = 2,
// the resulting row-1 will be (1 -8 -9).
func subtractRows(r1, r2 []float64, factor float64) {
	n := len(r1)
	for c := 0; c < n; c++ {
		r1[c] -= factor * r2[c]
	}
}
