package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func hypot(x, y float64) float64 {
	return math.Pow(x, 2) + math.Pow(y, 2)
}
func main() {
	r, _ := os.Open("something something")
	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
