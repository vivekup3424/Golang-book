package main

func main() {
	var x []int
	ch := make(chan int)
	go func() {
		x = make([]int, 10)
		ch <- 1
	}()
	go func() {
		x = make([]int, 1000)
		ch <- 1
	}()
	<-ch
	x[999] = 10
}
