package main

import (
	"main/pipeline/node"
	"fmt"
	"os"
	"bufio"
)

func main() {
	fileSort()
}

func fileSort() {
	const filename = "small.in"
	const n = 64

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close() //defer始终保持运行的

	p := node.RandomSource(n)

	//加一个bufio.NewWriter
	writer := bufio.NewWriter(file)
	node.WriteSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	p = node.ReaderSource(
		bufio.NewReader(file), -1)

	for v := range p {
		fmt.Println(v)
	}
}

func arraySort() {
	p1 := node.ArraySource(1, 3, 5, 8, 0)
	p1 = node.InMemSort(p1)

	p2 := node.ArraySource(6, 3, 7, 41)
	p2 = node.InMemSort(p2)
	//for {
	//	if num, ok := <-p; ok {
	//		fmt.Println(num)
	//	} else {
	//		break
	//	}
	//}

	p := node.Merge(p1, p2)
	for v := range p {
		//用range的话发送方一定要close
		fmt.Println(v)
	}
}
