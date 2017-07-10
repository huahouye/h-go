/*
数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。
因为数组的长度是固定的，因此在Go语言中很少直接使用数组。和数组对应的类型是
Slice（切片），它是可以增长和收缩动态序列，slice功能也更灵活，但是要理解slice工作原
理的话需要先理解数组。
数组的每个元素可以通过索引下标来访问，索引下标的范围是从0开始到数组长度减1的位
置。内置的len函数将返回数组中元素的个数。
*/
package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	main1()
	main2()
	main3()
}

func main1() {
	fmt.Println("main1()")
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2], r[2])
}

func main2() {
	/*
		在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始
		化值的个数来计算。
	*/
	fmt.Println("main2()")
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	// 数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长
	// 度必须是常量表达式，因为数组的长度需要在编译阶段确定
	// p := [3]int{1,2,3}
	// p = [4]int{1,2,3,4} // compile error: cannot assign [4]int to [3]int
	// 数组、slice、map和结构体字面值的写法都很相似。上面的形式是直接提供
	// 顺序初始化值序列，但是也可以指定一个索引和对应值列表的方式初始化，就像下面这样
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ¥"

	// 定义了一个含有100个元素的数组r，最后一个元素被初始化为-1，其它元素都是用0初始化
	r := [...]int{99: -1}
	fmt.Println(r[0], r[99])
}

// 数组比较
func main3() {
	fmt.Println("main3()")
	/*
		如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，这时候我
		们可以直接通过==比较运算符来比较两个数组，只有当两个数组的所有元素都是相等的时候
		数组才是相等的。不相等比较运算符!=遵循同样的规则
	*/
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	d := [3]int{1, 2}
	fmt.Println(d[0])
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}
