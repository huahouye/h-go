/*
Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两
种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实
部和虚部
*/
package main

import (
	"fmt"
)

func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"

	/*
	如果一个浮点数面值或一个十进制整数面值后面跟着一个i，例如3.141592i或2i，它将构成一
	个复数的虚部，复数的实部是0
	*/
	fmt.Println(1i * 1i) // "(-1+0i)", i^2 = -1
}
