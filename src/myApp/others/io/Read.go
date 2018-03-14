package io

import (
	"io"
	"strings"
	"fmt"
	"os"
	"bufio"
)

func readFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func readFromString() {
	data, _ := readFrom(strings.NewReader("Hello there !"), 12)
	fmt.Printf("%v\n", data)
}

func readFromStdin() {
	fmt.Println("please input from stdin :")
	data, _ := readFrom(os.Stdin, 11)
	fmt.Println(data)

	//需要go build Read.go ，然后运行Read.exe
}

func readFromFile() {
	file, _ := os.Open("src\\main\\io\\Read.go")
	defer file.Close()
	data, _ := readFrom(file, 12)
	fmt.Println(string(data))
}

//缓冲io
func test() {
	strReader := strings.NewReader("Hello God")
	bufReader := bufio.NewReader(strReader)
	data, _ := bufReader.Peek(7)
	fmt.Println(string(data))
	//缓冲了9个字符，Peek是偷看的意思
	fmt.Println(string(data), bufReader.Buffered())

	//读取到空格位置，所以最后只缓冲了3个字符God
	str, _ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered())

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello ") //写入到文件，包括屏幕
	fmt.Fprint(w, "Go ！")
	w.Flush()
}

func main() {
	//readFromString()
	//readFromStdin()
	//readFromFile()
	test()
}
