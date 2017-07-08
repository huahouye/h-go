/*
将一个表示整值的字符串，每隔三个字符插
入一个逗号分隔符，例如“12345”处理后成为“12,345”
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345"))
}

func comma(s string) string {
	n := len(s)
	// 如果输入字符串的长度小于或等于 3 的话，则不需要插入逗分隔符。
	if n <= 3 {
		return s
	}
	// 否则，comma 函数将在最后三个字符前位置将字符串切割为两个两个子串并插
	// 入逗号分隔符，然后通过递归调用自身来出前面的子串
	return comma(s[:n-3]) + "," + s[n-3:]
}
