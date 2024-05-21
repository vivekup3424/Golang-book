package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Pow(x, 2) + math.Pow(y, 2)
}
func main() {
	x := hypot(6, 6)
	fmt.Printf("Answer = %0.4f\n", x)
}
