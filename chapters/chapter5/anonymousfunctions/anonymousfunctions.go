/*
Anonymous Functions 匿名函数
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))

	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

// 更为重要的是，通过这种方式定义的函数可以访问完整的词法环境（lexical environment），
// 这意味着在函数中定义的内部函数可以引用该函数的变量，如下例所示
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
/*
函数squares返回另一个类型为 func() int 的函数。对squares的一次调用会生成一个局部变量
x并返回一个匿名函数。每次调用时匿名函数时，该函数都会先使x的值加1，再返回x的平
方。第二次调用squares时，会生成第二个x变量，并返回一个新的匿名函数。新匿名函数操
作的是第二个x变量。

squares的例子证明，函数值不仅仅是一串代码，还记录了状态。在squares中定义的匿名内
部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引
用。这就是函数值属于引用类型和函数值不可比较的原因。Go使用闭包（closures）技术实
现函数值，Go程序员也把函数值叫做闭包。
*/
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
