package node

import "sort"

// <-chan 只能输出的管道
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var a []int
		for v := range in {
			a = append(a, v)
		}

		sort.Ints(a)

		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}
