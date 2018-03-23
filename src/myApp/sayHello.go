package main

import (
	"fmt"
)

func main() {
	var map1 = map[string]int{}
	map2 := make(map[string]int)

	map1["key2"] = 50
	map2["key1"] = 5

	fmt.Println(map1["key2"])
	fmt.Println(map2["key1"])

}
