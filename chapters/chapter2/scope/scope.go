/*
作用域
在函数中词法域可以深度嵌套，因此内部的一个声明可能屏蔽外部的声明。还有许多语法块
是if或for等控制流语句构造的。下面的代码有三个不同的变量x，因为它们是定义在不同的词
法域（这个例子只是为了演示作用域规则，但不是好的编程风格）
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// 在 x[i]  和 x + 'A' - 'a'  声明语句的初始化的表达式中都引用了外部作用域声明的x变量
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}

	fmt.Println()
	main2()

	fmt.Println()
	main3()

	fmt.Println()
	main4()

	fmt.Println()
	main5()
}

func main2() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}

func main3() {
	if x := f(); x == 0 {
		fmt.Println("if ", x)
	} else if y := g(); x == y {
		// 第二个if语句嵌套在第一个内部，因此第一个if语句条件初始化词法域声明的变量在第二个if中也可以访问
		fmt.Println("else if", x, y)
	} else {
		fmt.Println("else", x, y)
	}
	// fmt.Println(x, y) // compile error: x and y are not visible here
}

func f() int {
	return 1
}

func g() int {
	return 1
}

func main4() {
	if f, err := os.Open("/tmp/test.txt"); err != nil {
		f.Close()
		fmt.Println("err %v", err)
		return
	}
	// f.ReadByte() // compile error: undefined f
	// f.Close() // compile error: undefined f
}

func main5() error {
	f, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println("err %v", err)
		return err
	}
	f.Close()
	return nil
}
