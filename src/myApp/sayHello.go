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


	utils.SayHello()
}
