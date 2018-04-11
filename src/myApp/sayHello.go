package main

import (
	"fmt"
)

//定义接口
type Sayer interface {
	SayHello()
}

//定义结构体Person
type Person struct {
	name string
}

func (person *Person) SayHello() {
	fmt.Println("Hello Person")
}

func main() {
	var say Sayer
	say = new(Person)

	//测试某一个值是否实现了某个接口，此处没有“*”
	if _, ok := say.(Sayer); ok {
		fmt.Println("Yes")
	}
}
