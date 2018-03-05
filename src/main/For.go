package main

import "fmt"

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}

	for i := 1; i < 5; i++ {
		fmt.Println(map1[i])
	}

	for i, v := range map1 {
		fmt.Printf("%d: %s\n", i, v)
	}

	for i, v := range "Go语言" {
		//println("%d: %c\n", i, v)
		fmt.Printf("%d: %c\n", i, v)
	}
}

/*
range子句包含一个或两个迭代变量（用于与迭代出的值绑定）、
特殊标记:=或=、
关键字range以及range表达式。
其中，range表达式的结果值的类型应该是能够被迭代的，
包括：字符串类型、数组类型、数组的指针类型、切片类型、字典类型和通道类型。
例如：

for i, v := range "Go语言" {
    fmt.Printf("%d: %c\n", i, v)
}

对于字符串类型的被迭代值来说，for语句每次会迭代出两个值。
第一个值代表第二个值在字符串中的索引，而第二个值则代表该字符串中的某一个字符。
迭代是以索引递增的顺序进行的。
例如，上面的for语句被执行后会在标准输出上打印出：

0: G
1: o
2: 语
5: 言

可以看到，这里迭代出的索引值并不是连续的。因为：
一个中文字符在经过UTF-8编码之后会表现为三个字节。

最后，我们来说一下break语句和continue语句。
它们都可以被放置在for语句的代码块中。
前者被执行时会使其所属的for语句的执行立即结束，
而后者被执行时会使当次迭代被中止（当次迭代的后续语句会被忽略）而直接进入到下一次迭代。
 */
