package main

func main() {

}

/*
select语句属于条件分支流程控制方法，不过它只能用于通道。
它可以包含若干条case语句，并根据条件选择其中的一个执行。
进一步说，select语句中的case关键字只能后跟用于通道的发送操作的表达式以及接收操作的表达式或语句。示例如下：

ch1 := make(chan int, 1)
ch2 := make(chan int, 1)
// 省略若干条语句
select {
case e1 := <-ch1:
    fmt.Printf("1th case is selected. e1=%v.\n", e1)
case e2 := <-ch2:
    fmt.Printf("2th case is selected. e2=%v.\n", e2)
default:
    fmt.Println("No data!")
}

如果该select语句被执行时通道ch1和ch2中都没有任何数据，
那么肯定只有default case会被执行。
但是，只要有一个通道在当时有数据就不会轮到default case执行了。
显然，对于包含通道接收操作的case来讲，其执行条件就是通道中存在数据（或者说通道未空）。
如果在当时有数据的通道多于一个，那么Go语言会通过一种伪随机的算法来决定哪一个case将被执行。

 */
