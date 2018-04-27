package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 0)
	fmt.Println("切片长度：", len(s), "；切片容量：", cap(s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println("切片长度：", len(s), "；切片容量：", cap(s))
	}

}
