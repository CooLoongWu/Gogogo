package main

import (
	"myApp/utils"
	"fmt"
)

func multi(str string) (string, bool) {
	if len(str) == 0 {
		return str, false
	}
	return str, true
}

func main() {
HERE:
	fmt.Println("我在29行")

	fmt.Println("我在32行")
	goto HERE

	utils.SayHello()
}
