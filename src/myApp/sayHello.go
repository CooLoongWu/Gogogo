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
