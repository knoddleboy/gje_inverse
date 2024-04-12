package gjeinverse

import (
	"strconv"
	"testing"
)

// TestParallelInvolution tests the correctness of the InverseParallel function by
// verifying if applying the inverse operation twice on a randomly generated matrix
// returns the original matrix itself. If the result does not match the original
// matrix, it fails the test.
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

// TestSerialDeterministic tests the deterministic behavior of the InverseParallel function.
// It computes randomly generated matrix's inverses serially. Then it verifies that the
// computed inverses are equal pairwise, ensuring the deterministic nature of the computation.
func TestParallelDeterministic(t *testing.T) {
	threads := 4

	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.Randomize()

		inv1, _ := InverseParallel(a.Copy(), threads)
		inv2, _ := InverseParallel(a.Copy(), threads)
		inv3, _ := InverseParallel(a.Copy(), threads)

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

// TestParallelConformsSerial tests if the InverseParallel function produces the same result
// as the InverseSerial function for randomly generated matrices of varying dimensions.
// It iterates through a set of test dimensions defined in TestDims and compares the inverses
// computed by both parallel and serial methods. If the results do not match, it fails the test.
func TestParallelConformsSerial(t *testing.T) {
	for p := 2; p <= 4; p++ {
		t.Run("p="+strconv.Itoa(p), func(t *testing.T) {
			for _, dim := range TestDims {
				a := NewMatrix(dim)
				a.Randomize()

				invPar, _ := InverseParallel(a.Copy(), p)
				invSer, _ := InverseSerial(a)

				if !invPar.Equals(invSer) {
					t.Fatal(ErrFailedToCompute)
				}
			}
		})
	}
}
