package main

import (
	"fmt"
)

func main() {
	num := make(chan int, 2)

	fmt.Println(num)
	num <- 1

	var value int
	value, _ = <-num
	fmt.Println(value)

	num <- 2
	value, _ = <-num
	fmt.Println(value)

	num <- 3
	value, _ = <-num
	fmt.Println(value)
}
