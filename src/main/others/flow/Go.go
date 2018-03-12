package flow

import (
	"fmt"
)

func main() {
	//go fmt.Println("Go !")
	////runtime.Gosched()
	//time.Sleep(1000)

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 3)
	go func() {
		fmt.Println("1")
		ch1 <- 1
	}()
	go func() {
		<-ch1
		fmt.Println("2")
		ch2 <- 2
		fmt.Printf("通道2输出结果%v", <-ch2)
	}()
	go func() {
		<-ch2
		fmt.Println("3")
		ch3 <- 3
	}()
	<-ch3

}

/*
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    go func() {
        fmt.Println("Go!")
        wg.Done()
    }()
    go func() {
        fmt.Println("Go!")
        wg.Done()
    }()
    go func() {
        fmt.Println("Go!")
        wg.Done()
    }()
    wg.Wait()
}
    sync.WaitGroup类型有三个方法可用——Add、Done和Wait。
Add会使其所属值的一个内置整数得到相应增加，
Done会使那个整数减1，
而Wait方法会使当前Goroutine（这里是运行main函数的那个Goroutine）阻塞直到那个整数为0。
这下你应该明白上面这个示例所采用的方法了。
我们在main函数中启用了三个Goroutine来封装三个go函数。
每个匿名函数的最后都调用了wg.Done方法，并以此表达当前的go函数会立即执行结束的情况。
当这三个go函数都调用过wg.Done函数之后，
处于main函数最后的那条wg.Wait()语句的阻塞作用将会消失，main函数的执行将立即结束。
 */
