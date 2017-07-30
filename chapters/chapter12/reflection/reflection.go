/*
Go语音提供了一种机制在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内
在操作，但是在编译时并不知道这些变量的具体类型。这种机制被称为反射。反射也可以让
我们将类型本身作为第一类的值类型处理。

两个至关重要的API是如何用反射机制的：一个是fmt包提供的字符串格式功能，另一个是类似
encoding/json和encoding/xml提供的针对特定协议的编解码功能。对于我们在4.6节中看到过
的text/template和html/template包，它们的实现也是依赖反射技术的。然后，反射是一个复杂
的内省技术，不应该随意使用，因此，尽管上面这些包内部都是用反射技术实现的，但是它
们自己的API都没有公开反射相关的接口。
*/
package reflection

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

/*
reflect.Type和reflect.Value
函数 reflect.TypeOf 接受任意的 interface{} 类型, 并返回对应动态类型的reflect.Type
*/
func main1() {
	// 其中 TypeOf(3) 调用将值 3 作为 interface{} 类型参数传入. 回到 7.5节 的将一个具体的值转为
	// 接口类型会有一个隐式的接口转换操作, 它会创建一个包含两个信息的接口值: 操作数的动态
	// 类型(这里是int)和它的动态的值
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
}

/*
因为 reflect.TypeOf 返回的是一个动态类型的接口值, 它总是返回具体的类型. 因此, 下面的代
码将打印 "*os.File" 而不是 "io.Writer". 稍后, 我们将看到 reflect.Type 是具有识别接口类型的
表达方式功能的.
*/
func main2() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"

	// 要注意的是 reflect.Type 接口是满足 fmt.Stringer 接口的. 因为打印动态类型值对于调试和日
	// 志是有帮助的, fmt.Printf 提供了一个简短的 %T 标志参数, 内部使用 reflect.TypeOf 的结果输出
	fmt.Printf("%T\n", 3) // "int"
}

/*
reflect 包中另一个重要的类型是 Value. 一个 reflect.Value 可以持有一个任意类型的值. 函数
reflect.ValueOf 接受任意的 interface{} 类型, 并返回对应动态类型的reflect.Value. 和
reflect.TypeOf 类似, reflect.ValueOf 返回的结果也是对于具体的类型, 但是 reflect.Value 也可
以持有一个接口值.
*/
func main3() {
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"

	// 调用 Value 的 Type 方法将返回具体类型所对应的 reflect.Type
	t := v.Type()           // a reflect.Type
	fmt.Println(t.String()) // "int"

	// 逆操作是调用 reflect.ValueOf 对应的 reflect.Value.Interface 方法. 它返回一个 interface{} 类型
	// 表示 reflect.Value 对应类型的具体值
	x := v.Interface()    // an interface{}
	i := x.(int)          // an int
	fmt.Printf("%d\n", i) // "3"
}
