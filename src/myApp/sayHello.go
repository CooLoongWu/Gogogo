package main

import (
	"fmt"
)

func main() {

	ite := map[int]string{1: "hello", 2: "hi"}
	fmt.Printf("Version A: Value of items: %v\n", ite)
	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = i
	}
	fmt.Printf("数据展示: %v\n", items)
	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("数据展示: %v\n", items2)

}
