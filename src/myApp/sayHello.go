package main

import (
	"fmt"
)

//或者简单的结构体如下
type T struct {
	a, b, _ int
	//“_”表示这个字段不会被用到
}

func main() {
	t := newT(1, 2, 3)
	fmt.Println(t)
}

func newT(a int, b int, c int) *T {
	if a < 0 {
		return nil
	}
	return &T{a, b, c}
}
