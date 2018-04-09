# 接口

## 定义
接口定义了一组方法，但是这些方法不包含代码，即没有被实现。 所以接口里也不能包含变量。 

例如下面的代码，定义了一个Sayer的接口类型，包含不带参数没有返回值的SayHello()方法和
不带参数有返回值的SayHi()方法以及一个带参数无返回值的SayGo(str string)方法：  
```go
type Sayer interface {
	SayHello()
	SayHi() string
	SayGo(str string)
}
```

## 使用
举例  
```go
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
```