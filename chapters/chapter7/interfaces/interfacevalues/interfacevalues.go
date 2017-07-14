/*
概念上讲一个接口的值，接口值，由两个部分组成，一个具体的类型和那个类型的值。它们
被称为接口的动态类型和动态值。对于像Go语言这种静态类型的语言，类型是编译期的概
念；因此一个类型不是一个值。在我们的概念模型中，一些提供每个类型信息的值被称为类
型描述符，比如类型的名称和方法。在一个接口值中，类型部分代表与之相关类型的描述
符。
*/
package interfacevalues

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

// 下面4个语句中，变量w得到了3个不同的值。（开始和最后的值是相同的）
func main1() {
	// 在Go语言中，变量总是被一个定义明确的值初始化，即使接口类型也不例外。对于一个接口
	// 用w==nil或者w!=nil来判读接口值是否为空
	// 调用一个空接口值上的任意方法都会产生panic
	// 的零值就是它的类型和值的部分都是nil
	var w io.Writer
	// w.Write([]byte("hello")) // panic: nil pointer dereference
	// 这个赋值过程调用了一个具体类型到接口类型的隐式转换，这和显式的使用
	// io.Writer(os.Stdout)是等价的
	// 这个接口值的动态类型被设为*os.Stdout指针的类型描述符，它的动态值持有os.Stdout
	w = os.Stdout
	w.Write([]byte("hello")) // "hello"
	// 效果和下面这个直接调用一样
	os.Stdout.Write([]byte("hello")) // "hello"
	// 现在动态类型是*bytes.Buffer并且动态值是一个指向新分配的缓冲区的指针
	w = new(bytes.Buffer)
	// 这个重置将它所有的部分都设为nil值，把变量w恢复到和它之前定义时相同的状态
	w = nil
}

func main2() {
	// 一个接口值可以持有任意大的动态值。例如，表示时间实例的time.Time类型，这个类型有几
	// 个对外不公开的字段。我们从它上面创建一个接口值
	var x interface{} = time.Now()
}

func main3() {
	/*
		接口值可以使用＝＝和！＝来进行比较。两个接口值相等仅当它们都是nil值或者它们的动态
		类型相同并且动态值也根据这个动态类型的＝＝操作相等。因为接口值是可比较的，所以它
		们可以用在map的键或者作为switch语句的操作数。
		然而，如果两个接口值的动态类型相同，但是这个动态类型是不可比较的（比如切片），将
		它们进行比较就会失败并且panic:
	*/
	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x) // panic: comparing uncomparable type []int

	/*
		考虑到这点，接口类型是非常与众不同的。其它类型要么是安全的可比较类型（如基本类型
		和指针）要么是完全不可比较的类型（如切片，映射类型，和函数），但是在比较接口值或
		者包含了接口值的聚合类型时，我们必须要意识到潜在的panic。同样的风险也存在于使用接
		口作为map的键或者switch的操作数。只能比较你非常确定它们的动态值是可比较类型的接口
		值。

		当我们处理错误或者调试的过程中，得知接口值的动态类型是非常有帮助的。所以我们使用
		fmt包的%T动作:
	*/
	// 在fmt包内部，使用反射来获取接口动态类型的名称
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil>"
	w = os.Stdout
	fmt.Printf("%T\n", w) // "*os.File"
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // "*bytes.Buffer"
}
