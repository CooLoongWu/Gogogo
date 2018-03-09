package node

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
)

// <-chan 只能输出的管道【调用该方法的人只能收取数据】
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//排序
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		//读取到内存中
		var a []int
		for v := range in {
			a = append(a, v)
		}

		//排序
		sort.Ints(a)

		//输出
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

//合并两个节点的数据【归并】
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {

	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2
	//获得 inputs[0..m) 和 inputs[m..end]
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))
}

//从文件读取数据
func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)

		bytesRead := 0

		for {
			n, err := reader.Read(buffer)

			bytesRead += n

			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}

			if err != nil ||
				(chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
