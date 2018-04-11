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

//定义结构体People
type People struct {
	name string
}

func (person *Person) SayHello() {
	fmt.Println("Hello Person")
}

func (people *People) SayHello() {
	fmt.Println("Hello People")
}

func main() {
	var say Sayer

	person := new(Person)

	say = person
	if _, ok := say.(*Person); ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	if _, ok := say.(*People); ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
