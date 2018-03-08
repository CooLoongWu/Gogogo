package main

import (
	"main/pipeline/node"
	"fmt"
)

func main() {
	p := node.ArraySource(1, 3, 5, 8, 0)
	p = node.InMemSort(p)
	//for {
	//	if num, ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}

	for v := range p {
		fmt.Println(v)
	}
}
