package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 1:
	case ch <- 2:
	}

	i := <-ch
	fmt.Println(i)
}
