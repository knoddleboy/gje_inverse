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

func measureParallelAverageWarmUp(dim, threads int) float64 {
	for i := 0; i < gjeinverse.SeqTests; i++ {
		a := gjeinverse.NewMatrix(dim)
		a.Randomize()
		gjeinverse.InverseSerial(a)
	}

	var total time.Duration

	for i := 0; i < gjeinverse.SeqTests; i++ {
		a := gjeinverse.NewMatrix(dim)
		a.Randomize()
		_, elapsed := gjeinverse.InverseParallel(a, threads)
		total += elapsed
	}

	avgTime := total / time.Duration(gjeinverse.SeqTests)
	fmt.Printf("(p)-%d,%d\tavg:\t%v\n", dim, threads, avgTime.Seconds())

	return avgTime.Seconds()
}

func main() {
	measureSerialAverageWarmUp(500)
	measureParallelAverageWarmUp(500, 4)
}
