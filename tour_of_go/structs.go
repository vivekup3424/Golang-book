package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{
		X: 1,
		Y: 2,
	}
	v2 := Vertex{3, 4}
	fmt.Println(v)
	fmt.Println(v2)
}
