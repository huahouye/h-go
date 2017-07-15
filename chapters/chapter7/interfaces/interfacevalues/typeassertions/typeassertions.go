/*
类型断言是一个使用在接口值上的操作。语法上它看起来像x.(T)被称为断言类型，这里x表示
一个接口的类型和T表示一个类型。一个类型断言检查它操作对象的动态类型是否和断言的类
型匹配。

这里有两种可能。
第一种，如果断言的类型T是一个具体类型，然后类型断言检查x的动态类
型是否和T相同。如果这个检查成功了，类型断言的结果是x的动态值，当然它的类型是T。换
句话说，具体类型的类型断言从它的操作对象中获得具体的值。如果检查失败，接下来这个
操作会抛出panic。
第二种，如果相反断言的类型T是一个接口类型，然后类型断言检查是否x的动态类型满足T。
如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口
值，但是结果有类型T。换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变
了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型和值的部分。
*/
package main

import (
	// "bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // success: f == os.Stdout
	fmt.Println(f == os.Stdout)
	// c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer

	main1()
	main2()
}

/*
Comma-ok断言的语法是：value, ok := element.(T)。element必须是接口类型的变量，
T是普通类型。如果断言失败，ok为false，否则ok为true并且value为变量的值。来看个例子：
*/
type Html []interface{}

func main1() {
	fmt.Println("main1()")
	html := make(Html, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	for index, element := range html {
		if value, ok := element.(string); ok {
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.([]byte); ok {
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		}
	}
}

/*
还有一种转换方式是switch测试。既然称之为switch测试，也就是说这种转换方式只能出现
在switch语句中。可以很轻松的将刚才用Comma-ok断言的例子换成由switch测试来实现
*/

func main2() {
	fmt.Println("main2()")
	html := make(Html, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = "style"
	html[4] = "head"
	for index, element := range html {
		switch value := element.(type) {
		case string:
			fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
		case []byte:
			fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
		default:
			fmt.Println("unknow type")
		}
	}
}
