package main

import (
	"fmt"
	"unicode/utf8"
	"github.com/fiam/gounidecode/unidecode"
	"strings"
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

	p := false
	fmt.Printf("|%T|%v|%#v|\n", p, p, p)

	fmt.Println("Hello World"[1])
	fmt.Println(".png"[1:])

	command := "set a 1"
	fmt.Println(command[4:])

	printf := func(format string, a ...interface{}){
		fmt.Println(fmt.Sprintf(format, a...))
	}

	printf("test %s, %d, %v", "string", 123, true)

	str = "Donec ante."
	fmt.Println(str[:utf8.RuneCountInString(str)-1])

	for i:=0;i<len(str);i++ {
		fmt.Println(str[i:i+1])
	}

	// transliteration
	str = "Количество байт в строках"
	fmt.Println(strings.Replace(strings.ToLower(unidecode.Unidecode(str)), " ", "_", -1))


}
