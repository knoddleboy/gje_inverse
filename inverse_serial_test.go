package gjeinverse

import (
	"testing"
)

// TestSerialInvolution tests the correctness of the InverseSerial function by verifying
// if applying the inverse operation twice on a randomly generated matrix returns
// the original matrix itself. It iterates through a set of test dimensions
// defined in TestDims and verifies the correctness for each dimension.
// If the result does not match the original matrix, it fails the test.
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
