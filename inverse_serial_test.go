package gjeinverse

import (
	"testing"
)

// TestSerialLargeDiagonal verifies the correctness of computing the inverse of
// a large diagonal matrix. It creates a diagonal matrix, fills it with values,
// computes its inverse serially, and checks using formula: aii = aii^-1
func TestSerialLargeDiagonal(t *testing.T) {
	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.FillDiagonal()

		inv, _ := InverseSerial(a)

		for i := 0; i < inv.Dim; i++ {
			for j := 0; j < inv.Dim; j++ {
				if i == j {
					if a.Data[i][i]-float64(i+1)*inv.Data[i][i] > 1e-8 {
						t.Fatal(i, ErrFailedToCompute)
					}
				} else {
					if inv.Data[i][j] != 0 {
						t.Fatal(i, ErrFailedToCompute)
					}
				}
			}
		}
	}
}

// TestSerialInvolution tests the correctness of the InverseSerial function by verifying
// if applying the inverse operation twice on a randomly generated matrix returns
// the original matrix itself. If the result does not match the original matrix,
// it fails the test.
func TestSerialInvolution(t *testing.T) {
	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.Randomize()

		inv, _ := InverseSerial(a.Copy())
		inv2, _ := InverseSerial(inv)

		if !inv2.Equals(a) {
			t.Fatal(ErrFailedToCompute)
		}
	}
}

// TestSerialDeterministic tests the deterministic behavior of the InverseSerial function.
// It computes randomly generated matrix's inverses serially. Then it verifies that the
// computed inverses are equal pairwise, ensuring the deterministic nature of the computation.
func TestSerialDeterministic(t *testing.T) {
	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.Randomize()

		inv1, _ := InverseSerial(a.Copy())
		inv2, _ := InverseSerial(a.Copy())
		inv3, _ := InverseSerial(a.Copy())

		if !inv1.Equals(inv2) {
			t.Fatal(ErrFailedToCompute)
		}
		if !inv2.Equals(inv3) {
			t.Fatal(ErrFailedToCompute)
		}
		if !inv1.Equals(inv3) {
			t.Fatal(ErrFailedToCompute)
		}
	}
}
