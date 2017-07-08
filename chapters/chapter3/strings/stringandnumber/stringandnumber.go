/*
除了字符串、字符、字节之间的转换，字符串和数值之间的转换也比较常见。由 strconv 包提
供这类转换功能
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 将一个整数转为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一个方法是用
	// strconv.Itoa(“整数到ASCII”)
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	// FormatInt和FormatUint函数可以用不同的进制来格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	// 如果要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无
	// 符号整数的ParseUint函数
	xx, xerr := strconv.Atoi("123")             // x is an int
	yy, yerr := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(xx, xerr, yy, yerr)
	/*
		ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int。在任何
		情况下，返回的结果y总是int64类型，你可以通过强制类型转换将它转为更小的整数类型
	*/
}
