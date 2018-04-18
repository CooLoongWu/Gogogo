package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string           //注意这里结构体变量需要 大写！！
	Age  int `json:"age"` //如果需要自定义json中的字段名可以在结构体变量后添加`json:"age"`类似的声明
}

func main() {
	////序列化
	//p := Person{
	//	Name: "LiMing", //赋值的时候也需要将结构体变量名称写出来
	//	Age:  8,
	//}
	//js, _ := json.Marshal(p) //使用json.Marshal()函数来序列化结构体实例
	//fmt.Printf("%s\n", js)
	//
	////反序列化
	//var person Person
	//json.Unmarshal(js, &person) //注意第二个参数需要取得对象的地址
	//fmt.Printf("%+v", person)

	var s []map[string]interface{}

	m1 := map[string]interface{}{"name": "LiMing", "age": 8}
	m2 := map[string]interface{}{"name": "LiLei", "age": 9}
	s = append(s, m1)
	s = append(s, m2)

	js, _ := json.Marshal(s)
	fmt.Printf("%s\n", js)

	var result []map[string]interface{}
	json.Unmarshal(js, &result)
	for key, value := range result {
		fmt.Println(key, value["age"], value["name"])
	}
}
