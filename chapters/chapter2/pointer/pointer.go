// 指针
package main

import (
	"fmt"
)

func main() {
	/*
		如果用“var x int”声明语句声明一个x变量，那么&x表达式（取x变量的内存地址）将产生
		一个指向该整数变量的指针，指针对应的数据类型是 *int，指针被称之为“指向int类型的指针”。
		如果指针名字为p，那么可以说“p指针指向变量x”，或者说“p指针保存了x变量的内存地址”。
		同时 *p  表达式对应p指针指向的变量的值。一般 *p  表达式读取指针指向的变量的值，这里
		为int类型的值，同时因为 *p  对应一个变量，所以该表达式也可以出现在赋值语句的左边，表
		示更新指针所指向的变量的值。
	*/
	i := 1
	p := &i         // p, of type *int, points to i
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to i = 2
	fmt.Println(i)  // "2"

	// 任何类型的指针的零值都是nil。如果 p != nil  测试为真，那么p是指向某个有效变量。指针
	// 之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

	// 在Go语言中，返回函数中局部变量的地址也是安全的。
	var pfp = pf()
	fmt.Println(*pfp)
	// 每次调用f函数都将返回不同的结果
	fmt.Println(pf() == pf()) // false

	// side effect: v is now 2
	// "3" (and v is 3)
	v := 100
	incr(&v)
	fmt.Println(incr(&v)) // 102
}

// 调用f函数时创建局
// 部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量
func pf() *int {
	v := 121
	return &v
}

/*
因为指针包含了一个变量的地址，因此如果将指针作为参数调用函数，那将可以在函数中通
过该指针来更新变量的值。例如下面这个例子就是通过指针来更新变量的值，然后返回更新
后的值，可用在一个表达式中（译注：这是对C语言中 ++v  操作的模拟，这里只是为了说
指针的用法，incr函数模拟的做法并不推荐）
*/
func incr(p *int) int { // 接受一个指向 int 类型变量的指针作为参数，如果 var x int那么 &x 就是符合的指针
	*p++ // 非常重要：只是增加 p 指向的变量的值，并不改变 p 指针！！！
	return *p
}
