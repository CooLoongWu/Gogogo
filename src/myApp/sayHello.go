package main

import "fmt"

func sum(values [] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	resultChan := make(chan int, 2)

	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)

	sum1, sum2 := <-resultChan, <-resultChan
	close(resultChan)

	fmt.Println("第一组计算结果：", sum1, "；第二组计算结果：", sum2)
	fmt.Println("计算结果:", sum1+sum2)
}
