/*
和方法"值"相关的还有方法表达式。当调用一个方法时，与调用一个普通的函数相比，我们必
须要用选择器(p.Distance)语法来指定方法的接收器。

当T是一个类型时，方法表达式可能会写作T.f或者(*T).f，会返回一个函数"值"，这种函数会将
其第一个参数用作接收器，所以可以用通常(译注：不写选择器)的方式来对其进行调用
*/
package methodvalues

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p *Point) Distance(p, q float64) float64 {
	return 0
}

func (p *Point) ScaleBy(p *Point, factor float64) float64 {
	return 0
}

func main1() {
	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
	// 译注：这个Distance实际上是指定了Point对象为接收器的一个方法func (p Point) Distance()，
	// 但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数，
	// 即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
	// 看起来本书中函数和方法的区别是指有没有接收器，而不像其他语言那样是指有没有返回值。
}

/*
当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。你可
以根据选择来调用接收器各不相同的方法。下面的例子，变量op代表Point类型的addition或者
subtraction方法，Path.TranslateBy方法会为其Path数组中的每一个Point来调用对应的方法
*/
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}
