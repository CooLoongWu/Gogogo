package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
	"bufio"
	"bytes"
)

func read(file io.Reader) ([]byte, error) {
	br := bufio.NewReader(file)
	var buf bytes.Buffer
	for {
		ba, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}
	}
	return buf.Bytes(), nil
}

func readFile(path string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullPath := filepath.Join(parentPath, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return read(file)
}

func main() {
	path := "src\\main\\Error.go"
	ba, err := readFile(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("The content of '%s':\n%s\n", path, ba)
}

/*
Go语言的函数可以一次返回多个结果。
这就为我们温和地报告错误提供了语言级别的支持。
实际上，这也是Go语言中处理错误的惯用法之一。我们先来回顾前一小节的例子：

func readFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return ioutil.ReadAll(file)
}
    函数readFile有两个结果声明。
第二个结果声明的类型是error。
error是Go语言内置的一个接口类型。它的声明是这样的：

type error interface {
    Error() string
}
    显然，只要一个类型的方法集合包含了名为Error、无参数声明且仅声明了一个string类型的结果的方法，
就相当于实现了error接口。

os.Open函数的第二个结果值就的类型就是这样的。
我们把它赋给了变量err。
也许你已经意识到，在Go语言中，函数与其调用方之间温和地传递错误的方法即是如此。

    在调用了os.Open函数并取得其结果之后，
我们判断err是否为nil。如果答案是肯定的，那么就直接把该错误（这里由err代表）返回给调用方。
这条if语句实际上是一条卫述语句。
这样的语句会检查流程中的某个步骤是否存在异常，
并在必要时中止流程并报告给上层的程序（这里是调用方）。

在Go语言的标准库以及很多第三方库中，我们经常可以看到这样的代码。
我们也建议大家在自己的程序中善用这样的卫述语句。

    现在我们把目光聚焦到readFile函数中的最后一条语句上。
这是一条return语句。
它把对ioutil.ReadAll函数的调用的结果直接作为readFile函数的结果返回了。
实际上，ioutil.ReadAll函数的结果声明列表与readFile的结果声明列表是一致的。
也就是说，它们声明的结果的数量、类型和顺序都是相同的。
因此，我们才能够做这种返回结果上的“嫁接”。
这又是一个Go语言编码中的惯用法。

    好了，在知晓怎样在传递错误之后，让我们来看看怎样创造错误。
没错，在很多时候，我们需要创造出错误（即error类型的值）并把它传递给上层程序。
这很简单。只需调用标准库代码包errors的New函数即可。
例如，我们只要在readFile函数的开始处加入下面这段代码就可以更快的在参数值无效时告知调用方：

if path == "" {
    return nil, errors.New("The parameter is invalid!")
}
    errors.New是一个很常用的函数。
在Go语言标准库的代码包中有很多由此函数创建出来的错误值，
比如os.ErrPermission、io.EOF等变量的值。
我们可以很方便地用操作符==来判断一个error类型的值与这些变量的值是否相等，
从而来确定错误的具体类别。就拿io.EOF来说，它代表了一个信号。
该信号用于通知数据读取方已无更多数据可读。
我们在得到这样一个错误的时候不应该把它看成一个真正的错误，而应该只去结束相应的读取操作。
请看下面的示例：

br := bufio.NewReader(file)
var buf bytes.Buffer
for {
    ba, isPrefix, err := br.ReadLine()
    if err != nil {
        if err == io.EOF {
            break
        }
        fmt.Printf("Error: %s\n", err)
        break
    }
    buf.Write(ba)
    if !isPrefix {
        buf.WriteByte('\n')
    }
}
    可以看到，这段代码使用到了前面示例中的变量file。
它的功能是把file代表的文件中的所有内容都读取到一个缓冲器（由变量buf代表）中。
请注意，该示例中的第6~8行代码。
如果判定err代表的错误值等于io.EOF的值（即它们是同一个值），
那么我们只需退出当前的循环以使读取操作结束即可。
 */
