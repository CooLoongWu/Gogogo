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