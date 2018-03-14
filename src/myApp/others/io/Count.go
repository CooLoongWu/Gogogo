package io

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var line int

	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		if !isPrefix {
			line++
		}
	}

	fmt.Println(line)
}

//需要go build Count.go生成Count.exe
//然后Count.exe Count.go来计算Count.go文件的代码行数
