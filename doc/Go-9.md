# 反射


## 定义  
反射是应用程序检查其所拥有的结构，尤其是类型的一种能力。这是元编程的一种形式。  

反射可以在运行时检查类型和变量，例如它的大小、方法和 动态 的调用这些方法。

reflect包可以在运行是自省类型、属性和方法，比如在一个变量上调用reflect.TypeOf()可以获取变量的正确类型，
如果变量是一个结构体类型，就可以通过Field来索引结构体的字段。  

反射可以从接口值反射到对象，也可以从对象反射回接口值。

## 使用
reflect.TypeOf()，reflect.ValueOf()的使用  

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = "hello"

	fmt.Println("变量str的type为:", reflect.TypeOf(str))

	v := reflect.ValueOf(str)
	fmt.Println("v.Type", v.Type())
	fmt.Println("v.Kind", v.Kind())
	fmt.Println("v.String", v.String())
	fmt.Println("v.Interface", v.Interface())

	s := v.Interface().(string)
	fmt.Println("变量str的value为:", s)
}

```

## 通过反射修改值
在上例中如果直接对v使用v.SetString()来修改v的值，因为reflect.ValueOf()函数通过传递一个str拷贝创建了v，所以v的改变
并不能更改原始的str。  

要使v的更改能作用到str，那么就必须传递str的地址。然后使用Elem()函数让v进行可设置。

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = "hello"

	v := reflect.ValueOf(str)
	fmt.Println("v.String", v.String())

	v = reflect.ValueOf(&str)   //注意这里是使用的str的地址
	v = v.Elem()                //需要Elem()函数使得v可以设置
	v.SetString("hi")
	fmt.Println("v.String", v.String())
	fmt.Println("str", str)
}

```

打印结果是：  
v.String hello  
v.String hi  
str hi