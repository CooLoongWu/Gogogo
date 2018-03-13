package server

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	//并发执行，go starts a goroutine
	for i := 0; i < 5; i++ {
		go printHelloWorld(i, ch)
	}

	for {
		msg := <-ch
		fmt.Println(msg)
	}

	//time.Sleep(time.Microsecond)
}

func printHelloWorld(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello world %d!\n", i)
	}

}
