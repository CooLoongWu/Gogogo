package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["str"] = "string"
	m["int"] = "int"

	fmt.Println("结果:", m["int"])
	fmt.Println("结果:", m["float"])

	delete(m, "int")
	fmt.Println("结果:", m["int"])

	var m1 map[string]int
	m1 = nil
	fmt.Println("结果：", m1["str"])
	m1["str"] = 1 //值为nil的map，相当于只读的map，这句写入操作会引发panic
	fmt.Println("结果：", m1["str"])
	delete(m1, "str")
}
