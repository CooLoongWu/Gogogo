package main

import (
	"myApp/utils"
	"fmt"
)

func main() {
	s := "hi"
	p := &s
	ss := *p

	fmt.Println("变量s在内存中的地址：", p, "；该内存地址上存放的是：", ss)
	fmt.Println("表达式s == *(&s)是否成立：", s == *(&s))
	utils.SayHello()
}
