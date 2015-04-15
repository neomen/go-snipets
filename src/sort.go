package main

import (
	"sort"
	"fmt"
)

type Item struct {
	Name		string
	Items		[]string
}

type Items []*Item

func (i Items) Len() int {
	return len(i)
}

func (i Items) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i Items) Less(a, b int) bool {
	return len(i[a].Items) < len(i[b].Items)
}

func main(){
	items := Items{
		&Item{"obj2", []string{"item1", "item2", "item3",}},
		&Item{"obj1", []string{"item1", "item2",}},
		&Item{"obj3", []string{"item1",}},
	}

	p := func(str string, items *Items){
		fmt.Println(str)
		for _, i := range *items {
			fmt.Println(i.Name)
		}
	}

	p("before:", &items)

	sort.Sort(items)

	p("after:", &items)
}
