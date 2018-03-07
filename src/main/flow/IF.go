package flow

func main() {
	var number int = 5
	if number += 4; 10 > number {
		number += 3
		println(number)
	} else if 10 < number {
		number -= 2
		println(number)
	}
	println(number)

	var number1 int
	if number1 := 4; 100 > number1 { //这里的number1并不是14行中声明的number1
		number1 += 3
	} else if 100 < number1 {
		number1 -= 2
	}
	println(number1) //这的number1才是14行声明的number1

	/*
	 这种写法有一个专有名词，叫做：标识符的重声明。
	实际上，只要对同一个标识符的两次声明各自所在的代码块之间存在包含的关系，
	就会形成对该标识符的重声明。具体到这里，
	第一次声明的number变量所在的是该if语句的外层代码块，
	而number := 4所声明的number变量所在的是该if语句的代表代码块。
	它们之间存在包含关系。因此对number的重声明就形成了。

    这种情况造成的结果就是，if语句内部对number的访问和赋值都只会涉及到第二次声明的那个number变量。
	这种现象也被叫做标识符的遮蔽。上述代码被执行完毕之后，
	第二次声明的number变量的值会是7，而第一次声明的number变量的值仍会是0。
	 */
}

/*
if number := 4; 100 > number {
    number += 3
} else if 100 < number {
    number -= 2
} else {
    fmt.Println("OK!")
}

这里的number := 4被叫做if语句的初始化子句。
它应被放置在if关键字和条件表达式之间，
并与前者由空格分隔、与后者由英文分号;分隔。

注意，我们在这里使用了短变量声明语句，
即：在声明变量number的同时为它赋值。
这意味着这里的number被视为一个新的变量。
它的作用域仅在这条i语句所代表的代码块中。
也可以说，变量number对于该if语句之外的代码来说是不可见的。
我们若要在该if语句以外使用number变量就会造成编译错误。
 */
