package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = "hello"

	v := reflect.ValueOf(str)
	fmt.Println("v.String", v.String())

	v = reflect.ValueOf(&str)
	v = v.Elem()
	v.SetString("hi")
	fmt.Println("v.String", v.String())
	fmt.Println("str", str)
}
