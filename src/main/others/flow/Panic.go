package flow

import (
	"errors"
	"fmt"
)

func innerFunc() {
	fmt.Println("Enter innerFunc")
	panic(errors.New("Occur a panic!"))
	fmt.Println("Quit innerFunc")
}

func outerFunc() {
	fmt.Println("Enter outerFunc")
	innerFunc()
	fmt.Println("Quit outerFunc")
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Fatal error: %s\n", p)
		}
	}()
	fmt.Println("Enter main")
	outerFunc()
	fmt.Println("Quit main")
}

/*
panic可被意译为运行时恐慌。因为它只有在程序运行的时候才会被“抛出来”。
并且，恐慌是会被扩散的。当有运行时恐慌发生时，它会被迅速地向调用栈的上层传递。
如果我们不显式地处理它的话，程序的运行瞬间就会被终止。这里有一个专有名词——程序崩溃。
内建函数panic可以让我们人为地产生一个运行时恐慌。
不过，这种致命错误是可以被恢复的。

在Go语言中，内建函数recover就可以做到这一点。

内建函数panic和recover是天生的一对。
前者用于产生运行时恐慌，而后者用于“恢复”它。
不过要注意，recover函数必须要在defer语句中调用才有效。
因为一旦有运行时恐慌发生，当前函数以及在调用栈上的所有代码都是失去对流程的控制权。
只有defer语句携带的函数中的代码才可能在运行时恐慌迅速向调用栈上层蔓延时“拦截到”它。

defer func() {
    if p := recover(); p != nil {
        fmt.Printf("Fatal error: %s\n", p)
    }
}()
    在这条defer语句中，我们调用了recover函数。
该函数会返回一个interface{}类型的值。
还记得吗？interface{}代表空接口。
Go语言中的任何类型都是它的实现类型。
我们把这个值赋给了变量p。
如果p不为nil，那么就说明当前确有运行时恐慌发生。
这时我们需根据情况做相应处理。
注意，一旦defer语句中的recover函数调用被执行了，运行时恐慌就会被恢复，
不论我们是否进行了后续处理。所以，我们一定不要只“拦截”不处理。

    我们下面来反观panic函数。
该函数可接受一个interface{}类型的值作为其参数。
也就是说，我们可以在调用panic函数的时候可以传入任何类型的值。
不过，我建议大家在这里只传入error类型的值。
这样它表达的语义才是精确的。
更重要的是，当我们调用recover函数来“恢复”由于调用panic函数而引发的运行时恐慌的时候，
得到的值正是调用后者时传给它的那个参数。
因此，有这样一个约定是很有必要的。

    总之，运行时恐慌代表程序运行过程中的致命错误。
我们只应该在必要的时候引发它。
人为引发运行时恐慌的方式是调用panic函数。
recover函数是我们常会用到的。
因为在通常情况下，我们肯定不想因为运行时恐慌的意外发生而使程序崩溃。
最后，在“恢复”运行时恐慌的时候，大家一定要注意处理措施的得当。

 */
