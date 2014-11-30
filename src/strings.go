package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "xyz"
	str2 := ""
	bool := true
	// len - количество байт в строке
	fmt.Printf("Количество байт в строках: %d %d\n",len(str), len(str2))

	// количество символов в строке
	count := utf8.RuneCountInString(str)
	fmt.Printf("Количество символов в строке %q: %d\n", str, count)

	fmt.Printf("bool: %t\n", bool)

	p := fasle
	fmt.Printf("|%T|%v|%#v|\n", p, p, p)
}
