/*
实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但
是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可
能会由于函数的间接引用被修改。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypot(3, 4)) // "5"

	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// 下面2个声明是等价的
func f(i, j, k int, s, t string)                 { /* ... */ }
func f2(i int, j int, k int, s string, t string) { /* ... */ }

// 4种方法声明拥有2个int型参数和1个int型返回值的函数.blank identifier
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return } // **
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

/*
你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义
了函数标识符。
package math
func Sin(x float64) float //implemented in assembly language
*/

// 准确的变量名可以传达函数返回值的含义。尤其在返回值的类型都相同时，就像下面这样
func Size(rect image.Rectangle) (width, height int)
func Split(path string) (dir, file string)
func HourMinSec(t time.Time) (hour, minute, second int)
