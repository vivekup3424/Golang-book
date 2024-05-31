package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		time.Sleep(0 * time.Second)
		fmt.Println("Today")
	case today + 1:
		time.Sleep(1 * time.Second)
		fmt.Println("Tomorrow")
	case today + 2:
		time.Sleep(2 * time.Second)
		fmt.Println("In two days")
	default:
		time.Sleep(5 * time.Second)
		fmt.Println("Too far away")
	}
}
