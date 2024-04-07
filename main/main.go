package main

import (
	"fmt"
	"time"

	gjeinverse "github.com/knoddleboy/gje_inverse"
)

func measureSerialAverageWarmUp(dim int) float64 {
	for i := 0; i < gjeinverse.SeqTests; i++ {
		a := gjeinverse.NewMatrix(dim)
		a.Randomize()
		gjeinverse.InverseSerial(a)
	}

	var total time.Duration

	for i := 0; i < gjeinverse.SeqTests; i++ {
		a := gjeinverse.NewMatrix(dim)
		a.Randomize()
		_, elapsed := gjeinverse.InverseSerial(a)
		total += elapsed
	}

	avgTime := total / time.Duration(gjeinverse.SeqTests)
	fmt.Printf("(s)-%d\tavg:\t%v\n", dim, avgTime.Seconds())

	return avgTime.Seconds()
}

func measureParallelAverageWarmUp(dim int) float64 {
	for i := 0; i < gjeinverse.SeqTests; i++ {
		a := gjeinverse.NewMatrix(dim)
		a.Randomize()
		gjeinverse.InverseSerial(a)
	}

	var total time.Duration

	for i := 0; i < gjeinverse.SeqTests; i++ {
		a := gjeinverse.NewMatrix(dim)
		a.Randomize()
		_, elapsed := gjeinverse.InverseParallel(a, 2)
		total += elapsed
	}

	avgTime := total / time.Duration(gjeinverse.SeqTests)
	fmt.Printf("(p)-%d\tavg:\t%v\n", dim, avgTime.Seconds())

	return avgTime.Seconds()
}

func main() {
	dim := 500
	ts := measureSerialAverageWarmUp(dim)
	tp := measureParallelAverageWarmUp(dim)

	fmt.Printf("Speed up:\t%.3f\n", ts/tp)

	// m := gjeinverse.NewMatrix(16)
	// m.Randomize()

	// inv, _, err := gjeinverse.InverseParallel(m.Copy(), 4)
	// if err != nil {
	// 	fmt.Println("error 1")
	// }

	// inv.Print()
	// fmt.Println()

	// inv2, _, err := gjeinverse.InverseParallel(inv.Copy(), 4)
	// if err != nil {
	// 	fmt.Println("error 2")
	// }

	// inv3, _, err := gjeinverse.InverseSerial(m.Copy())
	// if err != nil {
	// 	fmt.Println("error 2.2")
	// }

	// inv3.Print()

	// if !m.Equals(inv2) {
	// 	fmt.Println("error 3")
	// }
	// var total time.Duration = 0

	// for i := 0; i < 20; i++ {
	// 	m := gjeinverse.NewMatrix(1000)
	// 	m.Randomize()

	// 	gjeinverse.InverseParallel(m, 4)
	// }

	// for i := 0; i < 20; i++ {
	// 	m := gjeinverse.NewMatrix(1000)
	// 	m.Randomize()

	// 	// _, elapsed, _ := gjeinverse.InverseParallel(m, 4)
	// 	_, elapsed, _ := gjeinverse.InverseSerial(m)

	// 	total += elapsed
	// }

	// fmt.Printf("Avg:\t%v\n", total/20)

	// cpus := 4

	// a := gjeinverse.NewMatrix(16)
	// a.Randomize()

	// inv, _ := gjeinverse.InverseSerial(a.Copy())

	// inv.Print()

	// inv2, _ := gjeinverse.InverseParallel(a.Copy(), 4)

	// fmt.Println()
	// inv2.Print()

	// m := gjeinverse.NewMatrix(4)
	// m.Randomize()

	// m.Printf()

	// inv, _, _ := gjeinverse.InverseSerial(m.Copy())
	// inv, _, _ := gjeinverse.InverseParallel(m.Copy(), 4)

	// fmt.Printf("Parallel:\t%v\n", elapsed)
	// inv.Print()

	// fmt.Printf("Serial:\t%v\n", elapsed)
	// inv.Print()

	// _, elapsed_par, err := gjeinverse.InverseParallel(m.Copy(), 4)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Parallel:\t%v\n", elapsed_par)
	// inv.Print()

	// inv_ser, elapsed_ser, err := gjeinverse.InverseSerial(m.Copy())
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Serial:\t\t%v\n", elapsed_ser)

	// if inv.Equals(inv_ser) {
	// 	fmt.Println("-- correct --")
	// } else {
	// 	fmt.Println("-- error --")
	// }

	// speedup := float64(elapsed_ser) / float64(elapsed_par)
	// efficiency := speedup / float64(cpus)

	// fmt.Printf("Speed up:\t%.6f\n", speedup)
	// fmt.Printf("Efficiency:\t%.6f\n", efficiency)
}
