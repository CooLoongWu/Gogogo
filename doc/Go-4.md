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

### 声明格式
var name []type  
注意与数组的区别，切片不需要声明长度。  