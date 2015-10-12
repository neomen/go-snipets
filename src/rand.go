package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main() {
	fmt.Printf("Math rand num = %d\n", rand.Int())

	for i := 0; i<10; i++ {
		fmt.Println(rand.Int())
	}

	myrand := random(1, 99)
	fmt.Println(myrand)
}
