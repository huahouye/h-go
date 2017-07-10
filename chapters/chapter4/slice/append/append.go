/*
append函数对于理解slice底层是如何工作的非常重要，所以让我们仔细查看究竟是发生了什
么。下面是第一个版本的appendInt函数，专门用于处理[]int类型的slice
*/
package main

import (
	"fmt"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	/*
		每一次容量的变化都会导致重新分配内存和copy操作
		0 cap=1	[0]
		1 cap=2	[0 1]
		2 cap=4	[0 1 2]
		3 cap=4	[0 1 2 3]
		4 cap=8	[0 1 2 3 4]
		5 cap=8	[0 1 2 3 4 5]
		6 cap=8	[0 1 2 3 4 5 6]
		7 cap=8	[0 1 2 3 4 5 6 7]
		8 cap=16	[0 1 2 3 4 5 6 7 8]
		9 cap=16	[0 1 2 3 4 5 6 7 8 9]
	*/

	main1()

	var xx []int
	xx = appendInt2(xx, 1)
	xx = appendInt2(xx, 2, 3)
	fmt.Println(xx)
}
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
		// copy函数将返回成功复制的元素的个数（我们这里没有用到），等于两个
		// slice中较小的长度，所以我们不用担心覆盖会超出目标slice的范围。
	}
	z[len(x)] = y
	return z
}

/*
我们的appendInt函数每次只能向slice追加一个元素，但是内置的append函数则可以追加多个
元素，甚至追加一个slice。
*/
func main1() {
	fmt.Println("main1()")
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}

func appendInt2(x []int, y ...int) []int {
	fmt.Println("appendInt2(x []int, y ...int)")
	var z []int
	zlen := len(x) + len(y)
	// ... expand z to at least zlen
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
		// copy函数将返回成功复制的元素的个数（我们这里没有用到），等于两个
		// slice中较小的长度，所以我们不用担心覆盖会超出目标slice的范围。
	}
	copy(z[len(x):], y)
	return z
}
