package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Card struct {
	id     int
	Person []*Person
}

func main() {
	p1 := &Person{"李明", 8}
	p2 := &Person{"李雷", 9}

	card := Card{0, []*Person{p1, p2}}
	js1, _ := json.Marshal(p1)
	js2, _ := json.Marshal(p2)

	js3, _ := json.Marshal(card)

	fmt.Printf("%s\n", js1)
	fmt.Printf("%s\n", js2)
	fmt.Printf("%s\n", js3)

}
