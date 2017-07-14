/*
警告：一个包含nil指针的接口不是nil接口

一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。这个细微区别产生
了一个容易绊倒每个Go程序员的陷阱。
*/
package interfacevalues

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

/*
我们可能会预计当把变量debug设置为false时可以禁止对输出的收集，但是实际上在out.Write
方法调用时程序发生了panic

当main函数调用函数f时，它给f函数的out参数赋了一个*bytes.Buffer的空指针，所以out的动
态值是nil。然而，它的动态类型是*bytes.Buffer，意思就是out变量是一个包含空指针值的非
空接口（如图7.5），所以防御性检查out!=nil的结果依然是true。

问题在于尽管一个nil的*bytes.Buffer指针有实现这个接口的方法，它也不满足这个接口具体的
行为上的要求。特别是这个调用违反了(*bytes.Buffer).Write方法的接收者非空的隐含先觉条
件，所以将nil指针赋给这个接口是错误的。解决方案就是将main函数中的变量buf的类型改为
io.Writer，因此可以避免一开始就将一个不完全的值赋值给这个接口
*/
func main2() {
	var buf io.Writer // HERE
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}
