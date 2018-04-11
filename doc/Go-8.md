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

## 嵌套接口
嵌套接口就是在一个接口中包含了另一个接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
```go
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

```

## 类型断言
一个接口类型的变量中可以包含任何类型的值，所以我们需要来的检测他的类型，通常我们可以使用类型断言来测试在某个时候接口变量的值。  
类型断言可能是无效的，因为他不可能预见所有的可能性。如果转换在程序运行时失败会导致错误发生。

```go
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
	if _, ok := say.(*Person); ok {//say必须是接口变量，忘记“*”将会导致编译错误
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
```