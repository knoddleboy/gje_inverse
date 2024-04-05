package gjeinverse

import (
	"fmt"
	"testing"
	"time"
)

func TestInverseParallelIdentity(t *testing.T) {
	fmt.Println("Parallel - Identity")
	fmt.Println("dim\ttime, s")

	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.FillIdentity()

		inv, elapsed, err := InverseParallel(a, ThreadsTests)
		if err != nil {
			t.Fatal(err)
		}

		if !inv.IsIdentity() {
			t.Fatal(ErrFailedToCompute)
		}

		printDimTime(dim, elapsed.Seconds())
	}
}

func TestInverseParallelRandom(t *testing.T) {
	fmt.Println("Parallel - Random")
	fmt.Println("dim\ttime, s")

	for _, dim := range TestDims {
		a := NewMatrix(dim)
		a.Randomize()

		inv, elapsed, err := InverseParallel(a.Copy(), ThreadsTests)
		if err != nil {
			t.Fatal(err)
		}

		printDimTime(dim, elapsed.Seconds())

		inv2, _, err := InverseParallel(inv, ThreadsTests)
		if err != nil {
			t.Fatal(err)
		}

		if !a.Equals(inv2) {
			t.Fatal(ErrFailedToCompute)
		}
	}
}

func TestInverseParallelIdentitySeq(t *testing.T) {
	fmt.Println("Parallel - Identity (seq)")
	fmt.Println("dim\tavg time, s")

	for _, dim := range TestDims {
		for i := 0; i < SeqTests; i++ {
			a := NewMatrix(dim)
			a.FillIdentity()

			inv, _, err := InverseParallel(a, ThreadsTests)
			if err != nil {
				t.Fatal(err)
			}

			if !inv.IsIdentity() {
				t.Fatal(ErrFailedToCompute)
			}
		}

		var totalTime time.Duration

		for i := 0; i < SeqTests; i++ {
			a := NewMatrix(dim)
			a.FillIdentity()

			inv, elapsed, err := InverseParallel(a, ThreadsTests)
			if err != nil {
				t.Fatal(err)
			}

			totalTime += elapsed

			if !inv.IsIdentity() {
				t.Fatal(ErrFailedToCompute)
			}
		}

		avgTime := totalTime / time.Duration(SeqTests)
		printDimTime(dim, avgTime.Seconds())
	}
}

func TestInverseParallelRandomSeq(t *testing.T) {
	fmt.Println("Parallel - Random (seq)")
	fmt.Println("dim\tavg time, s")

	for _, dim := range TestDims {
		for i := 0; i < SeqTests; i++ {
			a := NewMatrix(dim)
			a.Randomize()

			inv, _, err := InverseParallel(a.Copy(), ThreadsTests)
			if err != nil {
				t.Fatal(err)
			}

			inv2, _, err := InverseParallel(inv, ThreadsTests)
			if err != nil {
				t.Fatal(err)
			}

			if !a.Equals(inv2) {
				t.Fatal(ErrFailedToCompute)
			}
		}

		var totalTime time.Duration

		for i := 0; i < SeqTests; i++ {
			a := NewMatrix(dim)
			a.Randomize()

			inv, elapsed, err := InverseParallel(a.Copy(), ThreadsTests)
			if err != nil {
				t.Fatal(err)
			}

			totalTime += elapsed

			inv2, _, err := InverseParallel(inv, ThreadsTests)
			if err != nil {
				t.Fatal(err)
			}

			if !a.Equals(inv2) {
				t.Fatal(ErrFailedToCompute)
			}
		}

		avgTime := totalTime / time.Duration(SeqTests)
		printDimTime(dim, avgTime.Seconds())
	}
}
