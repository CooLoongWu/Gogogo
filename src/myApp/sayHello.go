package main

import "fmt"

type Whisper interface {
	Say()
}

//定义接口
type Sayer interface {
	SayHello()
	Whisper
}

//定义结构体
type Person struct {
	name string
}

//实现接口的SayHello()
func (person *Person) SayHello() {
	fmt.Println("Hello " + person.name)
}

func (person *Person) Say() {
	fmt.Println("来自Whisper接口中的Say")
}

func main() {
	person := new(Person)
	person.name = "LiLei"
	person.SayHello()
	person.Say()
}
