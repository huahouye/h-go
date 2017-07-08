/*
bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、
byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要处理化，因为零
值也是有效的

bytes.Buffer类型有着很多实用的功能，我们在第七章讨论接口时将会涉及到，我们将看看如
何将它用作一个I/O的输入和输出对象，例如当做Fprintf的io.Writer输出对象，或者当作
io.Reader类型的输入源对象。
*/
// intsToString is like fmt.Sprint(values) but adds commas.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	// 当向bytes.Buffer添加任意字符的UTF8编码时，最好使用bytes.Buffer的WriteRune方法，但是
	// WriteByte方法对于写入类似'['和']'等ASCII字符则会更加有效
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
