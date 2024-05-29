package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func addTo(base int, vals ...int) ([]int, error) {
	var b = make([]int, len(vals))
	for i := 0; i < len(vals); i++ {
		b[i] = vals[i] + base
	}
	return b, nil
}
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return -1, -1, errors.New("Denominator is 0")
	} else {
		return numerator / denominator, numerator % denominator, nil
	}
}
func main() {
	fmt.Println(addTo(3, 1, 2, 3, 4, 5))
}
