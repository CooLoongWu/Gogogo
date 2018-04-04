package main

import "fmt"

type sayer interface {
	SayHello() string
}

type say struct {
	name string
}

func (ss *say) SayHello() string {
	return "Hello " + ss.name
}

func main() {
	s := new(say)
	s.name = "LiLei"

	ss := sayer(s)
	fmt.Println(ss.SayHello())
}
