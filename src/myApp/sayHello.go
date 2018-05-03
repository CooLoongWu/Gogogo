package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Num int       `json:"num"`
	Str string    `json:"str"`
	Sli []string  `json:"sli"`
	Arr [2]string `json:"arr"`
	Is  bool      `json:"is"`
	By  []byte    `json:"by"`
}

func main() {
	t := T{}
	js, _ := json.Marshal(t)
	fmt.Printf("%s", js)
}
