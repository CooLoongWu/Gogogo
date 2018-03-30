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
	//t := &T{1, 2, 4}

	//var _ T
	//t := T{1, 2, 4}

	//第三种
	t := T{a: 1}
	fmt.Println(t)
}
