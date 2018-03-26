# Map

## 概念 
map是一种特殊的数据结构，由键值对表示的无序集合。所以这个结构也成为关联数组或者字典。  

这是一种快速寻找值的理想结构：给定key，对应的value可以迅速定位。

## 声明

声明形式如下：var mapName map[keyType]valueType，例如：
```go
var map1 = map[string]int
```

在声明的时候不需要知道 map 的长度，map 是可以动态增长的。未初始化的 map 的值是 nil。

### 初始化
map 是 引用类型 的所以可以用 make 方法来分配内存。

```go
var map1 = map[string]int{}     //声明了一个空map1

map2 := make(map[string]int)    //使用make
```

### 键值对是否存在
如果map中不存在键key，那么获取他的值的时候就会得到一个类型的空值，所以我们就没有办法区分这个key到底是不存在
还是说这个key的值就是空值。  
解决办法：
```go
val, isExist := map[key]
```  
如果isExist是true，那么key存在，且val就是key对应的值。否则key就不存在。


### map类型的切片
如果要获取一个 map 类型的切片，我们必须使用两次 make() 函数，第一次分配切片，第二次分
配 切片中每个 map 元素，示例代码：
```go
items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = i
	}
	fmt.Printf("数据展示： %v\n", items)
```  

打印结果如下：  
数据展示: [map[1:0] map[1:1] map[1:2] map[1:3] map[1:4]]；
