package gjeinverse

import (
	"fmt"
	"testing"
	"time"
)

func TestInverseSerialIdentity(t *testing.T) {
	fmt.Println("Serial - Identity")
	fmt.Println("dim\ttime, s")

	for _, dim := range TestDims {
		m := NewMatrix(dim)
		m.FillIdentity()

		inv, elapsed, err := InverseSerial(m)
		if err != nil {
			t.Fatal(err)
		}

		if !inv.IsIdentity() {
			t.Fatal(ErrFailedToCompute)
		}

		printDimTime(dim, elapsed.Seconds())
	}
}

func TestInverseSerialRandom(t *testing.T) {
	fmt.Println("Serial - Random")
	fmt.Println("dim\ttime, s")

	for _, dim := range TestDims {
		m := NewMatrix(dim)
		m.Randomize()

		inv, elapsed, err := InverseSerial(m.Copy())
		if err != nil {
			t.Fatal(err)
		}

		printDimTime(dim, elapsed.Seconds())

		inv2, _, err := InverseSerial(inv)
		if err != nil {
			t.Fatal(err)
		}

		if !m.Equals(inv2) {
			t.Fatal(ErrFailedToCompute)
		}
	}
}

func TestInverseSerialIdentitySeq(t *testing.T) {
	fmt.Println("Serial - Identity (seq)")
	fmt.Println("dim\tavg time, s")

	for _, dim := range TestDims {
		for i := 0; i < SeqTests; i++ {
			m := NewMatrix(dim)
			m.FillIdentity()

			inv, _, err := InverseSerial(m)
			if err != nil {
				t.Fatal(err)
			}

			if !inv.IsIdentity() {
				t.Fatal(ErrFailedToCompute)
			}
		}

		var totalTime time.Duration

		for i := 0; i < SeqTests; i++ {
			m := NewMatrix(dim)
			m.FillIdentity()

			inv, elapsed, err := InverseSerial(m)
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

func TestInverseSerialRandomSeq(t *testing.T) {
	fmt.Println("Serial - Random (seq)")
	fmt.Println("dim\tavg time, s")

	for _, dim := range TestDims {
		for i := 0; i < SeqTests; i++ {
			m := NewMatrix(dim)
			m.Randomize()

			inv, _, err := InverseSerial(m.Copy())
			if err != nil {
				t.Fatal(err)
			}

			inv2, _, err := InverseSerial(inv)
			if err != nil {
				t.Fatal(err)
			}

			if !m.Equals(inv2) {
				t.Fatal(ErrFailedToCompute)
			}
		}

		var totalTime time.Duration

		for i := 0; i < SeqTests; i++ {
			m := NewMatrix(dim)
			m.Randomize()

			inv, elapsed, err := InverseSerial(m.Copy())
			if err != nil {
				t.Fatal(err)
			}

			totalTime += elapsed

			inv2, _, err := InverseSerial(inv)
			if err != nil {
				t.Fatal(err)
			}

			if !m.Equals(inv2) {
				t.Fatal(ErrFailedToCompute)
			}
		}

		avgTime := totalTime / time.Duration(SeqTests)
		printDimTime(dim, avgTime.Seconds())
	}
}
