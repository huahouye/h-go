/*
参数数量可变的函数称为为可变参数函数。典型的例子就是fmt.Printf和类似函数。Printf首先
接收一个的必备参数，之后接收任意个数的后续参数。
在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号“...”，这表示
该函数会接收任意数量的该类型参数。
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4))

	/*
	   在上面的代码中，调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一
	   个切片作为参数传给被调函数。如果原始参数已经是切片类型，我们该如何传递给sum？只需
	   在最后一个参数后加上省略符。下面的代码功能与上个例子中最后一条语句相同。
	*/
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))

	/*
		虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参数函
		数和以切片作为参数的函数是不同的。
	*/
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

/*
可变参数函数经常被用于格式化字符串。下面的errorf函数构造了一个以行号开头的，经过格
式化的错误信息。函数名的后缀f是一种通用的命名规范，代表该可变参数函数可以接收Printf
风格的格式化字符串。
*/
// interfac{}表示函数的最后一个参数可以接收任意类型
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
