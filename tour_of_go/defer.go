package main

import "fmt"

func main() {
	//a defer statement halts the execution of statement, till the
	//surrounding code has been executed
	defer fmt.Printf("World\n")
	fmt.Printf("Hello, ")
}
