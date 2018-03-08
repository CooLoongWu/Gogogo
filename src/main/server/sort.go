package main

import (
	"sort"
	"fmt"
)

func main() {
	a := []int{3, 56, 7, 0, 2, 6, 5, 6}
	sort.Ints(a)

	for i, v := range a {
		fmt.Println(i, v)
		//fmt.Println(i, v)
	}

	b := []string{"c", "a", "s", "d", "b"}
	sort.Strings(b)

	for i, v := range b {
		fmt.Println(i, v)
		//fmt.Println(i, v)
	}
}
