# Go-3 控制结构

## 3.1 if-else
*与java的区别：*  
省去了条件语句的括号

举例  
```go
var a string
if a == "" {
    fmt.Println("这是一个空字符串")
}
```

```go
a := "hello"
if len(a) == 0 {
    fmt.Println("这是一个空字符串")
} else {
    fmt.Println("这个字符串的长度是：", len(a))
}
```

```go
a := "hello"
if len(a) == 0 {
    fmt.Println("这是一个空字符串")
} else if len(a) == 5 {
    fmt.Println("这个字符串的长度是:5")
} else {
    fmt.Println("这个字符串的长度是：", len(a))
}
```

## 3.2 多返回值
*与java的区别：*   
一个函数可以返回多个返回值  

示例：  
```go
//os.Open()函数返回两个结果一个是文件，另一个是错误信息
f, err := os.Open(name)
if err !=nil {
    return err
}
```

定义一个返回多返回值的示例函数如下，函数multi()接受一个string作为参数，并返回string和boolean两个值：  
```go
//定义
func multi(str string) (string, bool) {
	if len(str) == 0 {
		return str, false
	}
	return str, true
}

//调用，无需使用的返回值也必须要用“_”表示出来
str, b := multi("")
_, b := multi("")
str, _ :=multi("")
```

## 3.3 switch
*与java的区别：*   
省去了条件语句的括号

## 3.4 for
*与java的区别：*   
省去了条件语句的括号  
用“for 条件语句 {}”来做循环

基本形式为：  
for 初始化语句; 条件语句; 修饰语句 {}

基于条件判断的迭代（类似java的 while 循环）：  
for 条件语句 {}   

无限循环：  
for {}    

基本形式的变种（类似java的foreach）：  
for index, val := range arr { } //表示循环一个数组中的数据

## 3.5 break、continue
同java

## 3.6 标签、goto
【不建议使用】
用一个全大写的字母并以“:”开头的形式称为标签，搭配goto可以使其返回执行标签下的代码，例如下面的例子会变成一个无限循环：  
  
```go
HERE:
	fmt.Println("我在第2行")
		
fmt.Println("我在4行")
goto HERE
```

## 3.7 Range
range 关键字用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素，直接看代码。

```go
package main

import (
	"fmt"
)

func main() {
	//range循环切片
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("计算结果：", sum)

	//range循环map
	maps := map[string]string{"a": "Angle", "b": "Bad"}
	for k, v := range maps {
		fmt.Println("map循环结果：", "key-value:", k, "-", v)
	}

	//range循环字符串，打印出来的值是unicode的值
	str := "Hello"
	for i, v := range str {
		fmt.Println("字符串循环结果：", "index-value:", i, "-", v)
	}
}

```