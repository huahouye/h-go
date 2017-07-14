/*
目前为止，我们看到的类型都是具体的类型。一个具体的类型可以准确的描述它所代表的值
并且展示出对类型本身的一些操作方式就像数字类型的算术操作，切片类型的索引、附加和
取范围操作。具体的类型还可以通过它的方法提供额外的行为操作。总的来说，当你拿到一
个具体的类型时你就知道它的本身是什么和你可以用它来做什么。

在Go语言中还存在着另外一种类型：接口类型。接口类型是一种抽象的类型。它不会暴露出
它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会展示出它们自
己的方法。也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是
可以通过它的方法来做什么。

在本书中，我们一直使用两个相似的函数来进行字符串的格式化：fmt.Printf它会把结果写到
标准输出和fmt.Sprintf它会把结果以字符串的形式返回。得益于使用接口，我们不必可悲的因
为返回结果在使用方式上的一些浅显不同就必需把格式化这个最困难的过程复制一份。实际
上，这两个函数都使用了另一个函数fmt.Fprintf来进行封装。fmt.Fprintf这个函数对它的计算
结果会被怎么使用是完全不知道的。
*/
package interfaceascontracts

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Fprintf的前缀F表示文件(File)也表明格式化输出结果应该被写入第一个参数提供的文件中。
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

// Printf函数中的第一个参数os.Stdout是*os.File类型
func Printf(format string, args ...interface{}) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}

// 在Sprintf函数中的第一个参数&buf是一个
// 指向可以写入字节的内存缓冲区，然而它 并不是一个文件类型尽管它在某种意义上和文件类
// 型相似
func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	Fprintf(&buf, format, args...)
	return buf.String()
}

/*
即使Fprintf函数中的第一个参数也不是一个文件类型。它是io.Writer类型这是一个接口类型

io.Writer类型定义了函数Fprintf和这个函数调用者之间的约定。一方面这个约定需要调用者提
供具体类型的值就像*os.File和*bytes.Buffer，这些类型都有一个特定签名和行为的Write的函
数。另一方面这个约定保证了Fprintf接受任何满足io.Writer接口的值都可以工作。Fprintf函数
可能没有假定写入的是一个文件或是一段内存，而是写入一个可以调用Write函数的值。
*/
type ByteCounter int

// 让我们通过一个新的类型来进行校验，下面*ByteCounter类型里的Write方法，仅仅在丢失写
// 向它的字节前统计它们的长度。
// 因为*ByteCounter满足io.Writer的约定，我们可以把它传入Fprintf函数中；Fprintf函数执行字
// 符串格式化的过程不会去关注ByteCounter正确的累加结果的长度。
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main1() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
