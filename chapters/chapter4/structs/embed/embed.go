/*
Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就
叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。下面的
代码中，Circle和Wheel各自都有一个匿名成员。我们可以说Point类型被嵌入到了Circle结构
体，同时Circle类型被嵌入到了Wheel结构体。
*/
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// 得意于匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径
	var w Wheel
	w.x = 8 // equivalent to w.Circle.Point.X = 8
	// w.Point.x = 8
	w.y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Printf("w %v\n", w)

	// 不幸的是，结构体字面值并没有简短表示匿名成员的语法， 因此下面的语句都不能编译通过
	// w = Wheel{8, 8, 5, 20} // compile error: unknown fields
	// w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
	// 结构体字面值必须遵循形状类型声明时的结构，所以我们只能用下面的两种语法，它们彼此
	// 是等价的
	w2 := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("w2 %v\n", w2)
	w3 := Wheel{
		Circle: Circle{
			Point: Point{
				x: 8,
				y: 8,
			},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("w3 %#v\n", w3)
}
