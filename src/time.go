package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	t := time.Now()
	fmt.Printf("The time is %s\n", t)

	var h time.Duration = 6
	t = t.Add(h*time.Hour)

	fmt.Printf("The time plus %d Hours -> %s\n", h, t.Local())
}
