# 接口

## 定义
接口定义了一组方法，但是这些方法不包含代码，即没有被实现。 所以接口里也不能包含变量。 

例如下面的代码，定义了一个say的接口类型：  
```go
type sayer interface {
	sayHello(name string) string
	sayHi(name string) string
}
```

## 使用
举例  
```go
package main

import "fmt"

//定义了一个接口sayer
type sayer interface {
	SayHello() string
}

//定义了一个结构体say
type say struct {
	name string
}

//
func (ss *say) SayHello() string {
	return "Hello " + ss.name
}

func main() {
	s := new(say)
	s.name = "LiLei"

	ss := sayer(s)
	fmt.Println(ss.SayHello())
}

```