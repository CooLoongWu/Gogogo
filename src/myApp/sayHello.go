package main

import (
	"fmt"
)

func main() {
	a := add()
	fmt.Println(a(1))
}

func add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
