/*
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。

例如，
*os.File类型实现了io.Reader，Writer，Closer，和ReadWriter接口。*bytes.Buffer实现了
Reader，Writer，和ReadWriter这些接口，但是它没有实现Closer接口因为它不具有Close方
法。Go的程序员经常会简要的把一个具体的类型描述成一个特定的接口类型。举个例子，
*bytes.Buffer是io.Writer；*os.Files是io.ReadWriter。
*/
package interfaces

import (
	"bytes"
	"io"
	"os"
	"time"
)

// 接口指定的规则非常简单：表达一个类型属于某个接口只要这个类型实现这个接口
func main1() {
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	// w = time.Second       // compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
	rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

	// 这个规则甚至适用于等式右边本身也是一个接口类型
	w = rwc // OK: io.ReadWriteCloser has Write method
	rwc = w // compile error: io.Writer lacks Close method
}

/*
就像信封封装和隐藏信件起来一样，接口类型封装和隐藏具体类型和它的值。即使具体类型
有其它的方法也只有接口类型暴露出来的方法会被调用到：
*/
func main2() {
	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello")) // OK: io.Writer has Write method
	// w.Close() // compile error: io.Writer lacks Close method
}

/*
一个有更多方法的接口类型，比如io.ReadWriter，和少一些方法的接口类型,例如io.Reader，
进行对比；更多方法的接口类型会告诉我们更多关于它的值持有的信息，并且对实现它的类
型要求更加严格。那么关于interface{}类型，它没有任何方法，请讲出哪些具体的类型实现了
它？
这看上去好像没有用，但实际上interface{}被称为空接口类型是不可或缺的。因为空接口类型
对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。
*/
func main3() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1, "two": 2}
	any = new(bytes.Buffer)
	/*
	对于创建的一个interface{}值持有一个boolean，float，string，map，pointer，或者任意其它
	的类型；我们当然不能直接对它持有的值做操作，因为interface{}没有任何方法。
	*/
}
