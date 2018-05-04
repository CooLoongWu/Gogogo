package main

import (
	"fmt"
	"crypto/md5"
)

func main() {
	str := "This is test ！"
	md := md5.New()
	md.Write([]byte(str))
	res := md.Sum([]byte(""))
	fmt.Printf("%x\n", res)
}
