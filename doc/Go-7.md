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

