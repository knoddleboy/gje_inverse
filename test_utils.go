package gjeinverse

import (
	"fmt"
)

// printDimTime prints the matrix dimension and time, taken for computing the inverse in a formatted way.
func printDimTime(dim int, time float64) {
	fmt.Printf("%d\t%.6f\n", dim, time)
}
