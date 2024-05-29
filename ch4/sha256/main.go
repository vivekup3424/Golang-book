package main

import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))

}
