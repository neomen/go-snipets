package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "275 / 45 R19 108 Y Стандартные"
	reg := regexp.MustCompile(`([0-9a-zA-ZСтандарт+])+`)
	tire := reg.FindAllString(str, -1)

	nums := make([]int, len(tire))
	nums[0], _ = strconv.Atoi(tire[0])
	nums[1], _ = strconv.Atoi(tire[1])
	nums[3], _ = strconv.Atoi(tire[3])

	reg = regexp.MustCompile(`([0-9])+`)
	rim := reg.FindAllString(tire[2], 1)
	nums[2], _ = strconv.Atoi(string(rim[0]))


	fmt.Printf("width: %d\n", nums[0])
	fmt.Printf("height: %d\n", nums[1])
	fmt.Printf("rim: %d\n", nums[2])
	fmt.Printf("load index: %d\n", nums[3])
	fmt.Printf("speed index: %s\n", tire[4])
	// standart
	fmt.Printf("standart: %s\n", tire[5])
	if strings.Contains(str, "Стандарт") {
		fmt.Printf("ok\n")
	} else {
		fmt.Printf("not ok\n")
	}

	//
	version := "A6 Quattro 2010->0"
//	version := "A6 Quattro 2010->2012"
//	version := "A6 Quattro"
	verArr := regexp.MustCompile(`([0-9]+)\-\>([0-9]+)`).FindAllString(version, -1)

	if(len(verArr)>0) {
		fmt.Println(verArr[0])
		str := strings.Split(verArr[0], "->")
		start, _ := strconv.Atoi(str[0])
		end, _ := strconv.Atoi(str[1])
		fmt.Printf("start: %d\n", start)
		fmt.Printf("end: %d\n", end)
	}

	//
	str = "275/45R19"
	reg = regexp.MustCompile(`([0-9])+`)
	tire = reg.FindAllString(str, -1)
	fmt.Println(tire)


	//
	// string pattern
	re := regexp.MustCompile(`api.contact.[0-9]+.[0-9]+.name.all`)
//	re2 := regexp.MustCompile(`api/contact/([0-9]+)/([0-9]+)/name/all`)

	fmt.Print("\n")
	fmt.Print(re.FindAllString("api.contact.10.1.name.all", -1))
	fmt.Print("\n")
//	fmt.Print("\n")
//	fmt.Printf("%q\n", re2.FindString("api/contact/10/1/name/all"))

	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)
	fmt.Println(regexp.MatchString("(api.contact.[0-9]+.[0-9]+.name.all):(get)", "api.contact.10.1.name.all:get"))
	fmt.Println(regexp.MatchString("(api.contact.[0-9]+.[0-9]+.name.all):(get)", "api.contact.10.1.name.all:post"))
	fmt.Println(regexp.MatchString("(api.contact.[0-9]+.[0-9]+.name.all):(get)", "api.contact.10.1.name.all"))
	fmt.Println(regexp.MatchString("(api.contact.[0-9]+.[0-9]+.name.all):(get)", "api.contact.10.1.name:get"))
	fmt.Println(regexp.MatchString("api.contact.[0-9]+.[0-9]+.name.all:get", "api.contact.10.1.name.all:get"))
	fmt.Println(regexp.MatchString("/api/post/[0-9//]*:get", "/api/post/:get"))
	fmt.Println(regexp.MatchString("/api/catalog/tire/[0-9]+/[0-9]+/[a-z_-]+:get", "/api/catalog/tire/25/1/-create_time:get"))

	// поиск и вывод
	s := `get published first 10 entries from 'markdown' category`
	fmt.Println(regexp.MatchString("get [a-z]+ [a-z]+ [0-9]+ entries from '[a-z]+' category", s))

	entries := regexp.MustCompile(`([0-9]+) entries`)
	num := entries.FindStringSubmatch(s)
	fmt.Println(num[1])

	category := regexp.MustCompile(`from '(.+)?' category`)
	cat := category.FindStringSubmatch(s)
	fmt.Println(cat[1])

	command := strings.Split(regexp.MustCompile(`(.+)?`).FindAllString(s, -1)[0], " ")
	fmt.Println(command)

	s2 := `get by id 1,2,3,4,5`
	fmt.Println(regexp.MatchString("get by id [0-9]+", s2))
	ids := regexp.MustCompile(`([0-9]+)`).FindAllString(s2, -1)
	fmt.Println(ids[3])
}
