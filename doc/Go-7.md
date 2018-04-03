# 结构

## 定义
结构体的定义一般方式如下：
```
type T struct{
	name string
	...
}

//或者简单的结构体如下
type T struct{
    a, b, _ int
    //“_”表示这个字段不会被用到
}

//声明一个结构体变量后可以在方法中使用以下方式赋值
var t T
t.a = 3
t.b = 4
```

结构体的字段可以是任何类型，甚至是结构体本身，也可以是函数或者接口。  

## new
使用new函数给一个新的结构体变量分配内存，他返回指向已分配内存的指针。
```go
var t *T = new(T)  

t := new(T)
//上面两条语句相同，变量t是一个指向T的指针
//！！如果T是一个结构体类型，那么表达式 new(T)和&T{}是等价的。


var t T
//上面代码也会给t分配内存，并初始化为零值，但是t是类型T。
```

## 初始化
```go
package main

import (
	"fmt"
)

type T struct {
	a, b, _ int
	//“_”表示这个字段不会被用到
}

func main() {
	//第一种
	t := &T{1, 2, 4}
	
	//第二种
	//var _ T
	//t := T{1, 2, 4}
	//上面两种方式都是  初始化一个结构体实例
	
	//第三种，只给指定字段赋值，其他字段为零值
	//t := T{a:1}
	fmt.Println(t)
}

```
### 工厂方法构建结构体实例
在 Go 语言中常常像下面这样在工厂方法里使用初始化来简便的实现构造子工厂。  
如果T是一个结构体类型，那么表达式 new(T)和&T{}是等价的。

```go
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
	t := newT(1, 2, 3)
	fmt.Println(t)
}

func newT(a int, b int, c int) *T {
	if a < 0 {
		return nil
	}
	return &T{a, b, c}
}

```

### 匿名字段和内嵌结构体
结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须
的，此时类型也就是字段的名字。匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。

请看下面代码示例：  
```go
package main

import (
	"fmt"
)

type inner struct {
	a int
	b int
}

type outer struct {
	c int
	string
	inner //上面的结构体，在这里也就是内嵌结构体
}

func main() {
	outerS1 := new(outer)
	outerS1.a = 1
	outerS1.b = 2
	outerS1.c = 3
	outerS1.string = "hello"
	fmt.Println("第一种方式输出：", outerS1)

	outerS2 := outer{3, "hello", inner{1, 2}}
	fmt.Println("第二种方式输出：", outerS2)
}

//最后结果打印：
//第一种方式输出： &{3 hello {1 2}}
//第二种方式输出： {3 hello {1 2}}
```

### 内嵌结构体字段冲突

结构体A和B中都有变量a，C、D都内嵌A、B，例如下面的代码：  
```go
package main

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

	d := new(D)
	d.a = 1		//没问题，因为D中本身就有一个a
}

```
所以：  
1、外层的名字会覆盖内层的名字，这就提供了一种重载字段或者方法的方式  
2、当同一名字在同一级别出现多次，如果使用该名字则必须明确表示需要具体结构体中的名字
