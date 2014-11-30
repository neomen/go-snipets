package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	p := []int{2,3,5,7,11,13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	//////////////////
	//Slicing slices
	//////////////////
	fmt.Println("Slicing slices")
	fmt.Println("p[1:4] ==", p[1:4])
	// missing low index implies 0
	fmt.Println("missing low index implies 0")
	fmt.Println("p[:3] ==", p[:3])
	// missing high index implies len(s)
	fmt.Println("missing high index implies len(s)")
	fmt.Println("p[4:] ==", p[4:])

	//////////////////
	//Making slices
	//////////////////
	fmt.Println("Making slices")
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	//////////////////
	//Nil slices
	//////////////////
	fmt.Println("Nil slices")
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
