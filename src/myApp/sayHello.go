package main

import "fmt"

//定义接口
type Sayer interface {
	SayHello()
	SayHi() string
	SayGo(str string)
}

//定义结构体
type Person struct {
	name string
}

//实现接口的SayHello()
func (person *Person) SayHello() {
	fmt.Println("Hello " + person.name)
}

//实现接口的SayHi()
func (person *Person) SayHi() string {
	msg := "Hi, " + person.name
	fmt.Println(msg)
	return msg
}

//实现接口的SayHi()
func (person *Person) SayGo(str string) {
	msg := "Go, " + person.name + ". " + str
	fmt.Println(msg)
}

func main() {
	person1 := new(Person)
	person1.name = "LiLei"
	//如果使用var的方式来定义sayer，那么结构体Person需要实现全部的接口
	//你可以试试注释掉SayGo的实现方法，看看会报什么错误
	var sayer1 Sayer
	sayer1 = person1
	sayer1.SayHello()

	//不使用var的方式来定义sayer，则无需实现全部的接口
	person2 := new(Person)
	person2.name = "LiMing"
	sayer2 := person2
	sayer2.SayHi()

	person3 := new(Person)
	person3.name = "LiSi"
	person3.SayGo("NoThing")
}
