package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	p := fmt.Println
	p("Welcome to the playground!")

	t := time.Now()
	fmt.Printf("The time is %s\n", t)

	var h time.Duration = 6
	t = t.Add(h*time.Hour)

	fmt.Printf("The time plus %d Hours -> %s\n", h, t.Local())

	// сравнение времени
	t2 := t.Add(time.Second * 60)
	p(t.Sub(t2).String())

	if strings.Contains(t.Sub(t2).String(), "-") {
		p("-")
	} else {
		p("+")
	}

	// время из строки
	str := "2015-04-08 21:27:52 +0700 +0700"
	t3, _ := time.Parse(time.RFC1123Z, str)
	p(t3)
}
