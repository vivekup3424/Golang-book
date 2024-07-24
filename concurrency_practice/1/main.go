package main

import (
	"fmt"
	"time"
)

func main() {
	//Kakashi notes down the duration of attack
	start := time.Now()
	defer func() {
		fmt.Println("Attack duration: ", time.Since(start))
	}()
	smokeSignal := make(chan bool)
	evilNinjas := []string{"Orochimaru", "Sasuke", "Tobi", "Kisame", "Uchiha Madara", "Sakura"}
	for i := range evilNinjas {
		attack(evilNinjas[i],len(evilNinjas)/4,smokeSignal)
	}
	<-smokeSignal
}
func attack(target string, secondsTaken int, smokeSignal chan bool) {
	fmt.Println("Throwing shuriken at : ", target)
	time.Sleep(time.Duration(secondsTaken) * time.Second) //time taken for shuriken to hit the target
	smokeSignal <- true
}
