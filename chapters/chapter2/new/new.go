/* 另一个创建变量的方法是调用用内建的new函数。表达式new(T)将创建一个T类型的匿名变
量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为 *T
*/
// new函数使用常见相对比较少，因为对应结构体来说，可以直接用字面量语法创建新变量的方法会更灵活
package main

import (
	"fmt"
)

func main() {
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // 0
	*p = 2
	fmt.Println(*p) // 2

	// 每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的
	pp := new(int)
	qq := new(int)
	fmt.Println(pp == qq)
}

// 两个newInt函数有着相同的行为
func newInt1() *int {
	return new(int)
}
func newInt2() *int {
	var v int
	return &v
}
