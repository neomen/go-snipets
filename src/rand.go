package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("Math rand num = %d\n", rand.Int())

	for i := 0; i<10; i++ {
		fmt.Println(rand.Int())
	}
}
