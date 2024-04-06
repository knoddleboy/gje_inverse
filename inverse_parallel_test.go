package gjeinverse

import (
	"testing"
)

// TestParallelInvolution tests the correctness of the InverseParallel function by verifying
// if applying the inverse operation twice on a randomly generated matrix returns
// the original matrix itself. It iterates through a set of test dimensions
// defined in TestDims and verifies the correctness for each dimension.
// If the result does not match the original matrix, it fails the test.
func TestParallelInvolution(t *testing.T) {
	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.Randomize()

		inv, _ := InverseParallel(a.Copy(), 4)
		inv2, _ := InverseParallel(inv, 4)

		if !inv2.Equals(a) {
			t.Fatal(ErrFailedToCompute)
		}
	}
}

// TestParallelConformsSerial tests if the InverseParallel function produces the same result
// as the InverseSerial function for randomly generated matrices of varying dimensions.
// It iterates through a set of test dimensions defined in TestDims and compares the inverses
// computed by both parallel and serial methods. If the results do not match, it fails the test.
func TestParallelConformsSerial(t *testing.T) {
	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.Randomize()

		invPar, _ := InverseParallel(a.Copy(), 4)
		invSer, _ := InverseSerial(a)

		if !invPar.Equals(invSer) {
			t.Fatal(ErrFailedToCompute)
		}
	}
}
