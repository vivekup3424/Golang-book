package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	YEN
	INR
	PKI
	CZH
)

func main() {
	var q [3]int = [3]int{1, 2, 3}
	r := [2]int{5, 6}
	fmt.Println(q)
	fmt.Println(r)
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, b == c, a == c)
	d := [...]int{1, 3, 4}

	//enums
}
