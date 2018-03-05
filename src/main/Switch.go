package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4)%2 ]; interface{}(v).(type) {
	case int32:
		fmt.Printf("Case A.")
	case byte:
		fmt.Printf("Case B.")
	default:
		fmt.Println("Unknown!")
	}
}

/*

在此类switch语句中，每个case会携带一个表达式。
与if语句中的条件表达式不同，这里的case表达式的结果类型并不一定是bool。
不过，它们的结果类型需要与switch表达式的结果类型一致。
所谓switch表达式是指switch语句中要被判定的那个表达式。
switch语句会依据该表达式的结果与各个case表达式的结果是否相同来决定执行哪个分支。
示例如下：

var name string
// 省略若干条语句
switch name {
case "Golang":
    fmt.Println("A programming language from Google.")
case "Rust":
    fmt.Println("A programming language from Mozilla.")
default:
    fmt.Println("Unknown!")
}

另外，与if语句一样，switch语句还可以包含初始化子句，且其出现位置和写法也如出一辙。


fallthrough：它既是一个关键字，又可以代表一条语句。
fallthrough语句可被包含在表达式switch语句中的case语句中。
它的作用是使控制权流转到下一个case。
不过要注意，fallthrough语句仅能作为case语句中的最后一条语句出现。
并且，包含它的case语句不能是其所属switch语句的最后一条case语句。
 */
