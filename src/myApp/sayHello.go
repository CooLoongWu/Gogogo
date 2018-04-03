package main

import "fmt"

type A struct {
	a int
	b int
}

type B struct {
	a int
}

type C struct {
	A
	B
}

type D struct {
	a int
	A
	B
}

func main() {
	c := new(C)
	//c.a = 1  //错误，因为这里无法判断是A中的a还是B中的a
	c.A.a = 1
	c.b = 2

	cc := C{}
	fmt.Println("值", cc.b)

	d := new(D)
	d.a = 1 //没问题，因为D中本身就有一个a
}
