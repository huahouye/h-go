/*
通过嵌入结构体来扩展类型
*/
package structembedding

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main1() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
}

/*
读者如果对基于类来实现面向对象的语言比较熟悉的话，可能会倾向于将Point看作一个基
类，而ColoredPoint看作其子类或者继承类，或者将ColoredPoint看作"is a" Point类型。
*/
func main2() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
}

/*
一个ColoredPoint并不是一个Point，但他"has a"Point，并且它有从Point类里引入的Distance
和ScaleBy方法。如果你喜欢从实现的角度来考虑问题，内嵌字段会指导编译器去生成额外的
包装方法来委托已经声明好的方法，和下面的形式是等价的
*/
func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}
func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

// 一个struct类型也可能会有多个匿名字段。我们将ColoredPoint定义为下面这样
type ColoredPoint struct {
	Point
	color.RGBA
}

/*
然后这种类型的值便会拥有Point和RGBA类型的所有方法，以及直接定义在ColoredPoint中的
方法。当编译器解析一个选择器到方法时，比如p.ScaleBy，它会首先去找直接定义在这个类
型里的ScaleBy方法，然后找被ColoredPoint的内嵌字段们引入的方法，然后去找Point和
RGBA的内嵌字段引入的方法，然后一直递归向下找。如果选择器有二义性的话编译器会报
错，比如你在同一级里有两个同名的方法。
*/
