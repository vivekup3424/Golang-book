package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	data := make(map[int]int)
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data[i] = i
		}(i)
	}
	wg.Wait()
	fmt.Println(data)
}
