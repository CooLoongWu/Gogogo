# Go-4 数组与切片

## 3.1 数组

### 概念
数组是具有相同 **唯一类型** 的一组已编号且长度固定的数据项序列【go支持多维数组】

### 声明格式
var name [length]type  
例如：  
```go
var name [5]string  
var age [2]int
```

### 初始化格式
直接初始化一个数组的时候可以这么写：  
```go
var arr [5]int
arr = [5]int{2, 3, 4}

//这样初始化的话，在索引3，4的位置上值默认初始化为0
arr1 := [5]int{1, 2, 3}

//这样初始化的话，在索引2，4的位置被赋予了实际值，其他元素都为空字符串
var arr2 = [5]string{2: "hello", 4: "hi"}

//当数组元素写全时，也可以使用...来省略数字
var arr3 = [...]int{2, 3, 5, 6}
```

还可以使用for循环的方式来进行初始化：  
```go
var arr [5]int
for i := 0; i < len(arr); i++ {
    arr[i] = i
}
```

### 使用格式
可以使用for循环为数组赋值或者打印数组  

```go
var arr [5]int
for i := 0; i < len(arr); i++ {
    arr[i] = i * 2
}

for i := 0; i < len(arr); i++ {
    fmt.Printf("索引位置%d是%d\n", i, arr[i])
}

for i, v := range arr {
    fmt.Printf("索引位置%d是%d\n", i, v)
}
```

## 切片
### 概念
切片（slice）是对数组一个连续片段的引用，所以切片是引用类型。  
和数组一样，切片通常也是一维的，但是也可以由一维组合成高维。

### 声明格式
var name []type  
注意与数组的区别，切片不需要声明长度。  

### 初始化格式
 var slice []type = arr[start:end]  
 这里slice是数组arr从start索引到end-1索引之间的元素构成的子集。  
 
 技巧：  
 var slice []type = arr[0:len(arr)]的简写就是  var slice []type = arr[:]  
 slice就等于完整的arr数组，另一种写法是：slice = &arr。  
 
 arr[2:]和arr[2:len(arr)]意义相同。  
 arr[:3]和arr[0:3]意义相同。
 
 
 注：  
 1、一个切片 s 可以这样扩展到它的大小上限： s = s[:cap(s)]   
 2、去掉 slice 的最后一个元素，只要 slice = slice[:len(slice)-1]  
 
 当相关数组还没有定义时，我们可以使用 make() 函数来创建一个切片 同时创建好相关数组：
 var slice []type = make([]type, len) ，这里 len 是数组的长度并且也是 slice 
 的初始长度。当然如果slice := make([]type, len, cap) ，这样的话就表示该切片不占用整个数组，
 而是只包含数组的前len个元素。


### 使用方式
类似数组，请看下面的代码：
```go
package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 2)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)
}

//猜猜输出结果是什么：
//结果是：[0 0 1 2]，因为切片通过make()函数初始化了，也就是说s包含的两个值都默认初始化了，后面append()才从第三个开始添加
```
  
再看一段代码：  
```go
package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 0)
	fmt.Println("切片长度：", len(s), "；切片容量：", cap(s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println("切片长度：", len(s), "；切片容量：", cap(s))
	}

}
/*
打印结果如下所示：

切片长度： 0 ；切片容量： 0
切片长度： 1 ；切片容量： 1
切片长度： 2 ；切片容量： 2
切片长度： 3 ；切片容量： 4
切片长度： 4 ；切片容量： 4
切片长度： 5 ；切片容量： 8
切片长度： 6 ；切片容量： 8
切片长度： 7 ；切片容量： 8
切片长度： 8 ；切片容量： 8
切片长度： 9 ；切片容量： 16
切片长度： 10 ；切片容量： 16
*/
```
可以看到在元素数量比较少时，当切片长度超过切片容量后，容量增长总是以当前容量的一倍增长。  
当元素数量超过1000后，每次增长则变为25%。
### 切片重组
将切片扩展 1 位可以这么做： slice = slice[0:len(slice)+1]
