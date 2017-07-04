/*
元组赋值是另一种形式的赋值语句，它允许同时更新多个变量的值。在赋值之前，赋值语句
右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。这对于处理有些同
时出现在元组赋值语句左右两边的变量很有帮助
*/
package main

import (
	"fmt"
)

func main() {
	// 最大公约数（GCD）greatest common divisor
	v := gcd(50, 30)
	fmt.Println(v)
	v2 := gcd(30, 50)
	fmt.Println(v2)

	f := fib(5)
	fmt.Println(f)
}

/*
	欧几里德的GCD是最早的非平凡算法
*/
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

/*
	斐波纳契数列（Fibonacci）
*/
func fib(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}
