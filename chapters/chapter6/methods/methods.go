/*
OOP编程的第一方面，我们会向你展示如何有效地定义和使用方法。我们会覆盖
到OOP编程的两个关键点，封装和组合。
*/
package methods

import (
	"fmt"
	"math"
)

/*
在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附
加到这种类型上，即相当于为这种类型定义了一个独占的方法。
*/

type Point struct {
	x, y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

/*
上面的代码里那个附加的参数p，叫做方法的接收器(receiver)，早期的面向对象语言留下的遗
产将调用一个方法称为“向一个对象发送消息”。

在Go语言中，我们并不会像其它语言那样用this或者self作为接收器；我们可以任意的选择接
收器的名字。由于接收器的名字经常会被使用到，所以保持其在方法间传递时的一致性和简
短性是不错的主意。这里的建议是可以使用其类型的第一个字母，比如这里使用了Point的首
字母p。
*/

func main1() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call

	/*
	   这种p.Distance的表达式叫做选择器，因为他会选择合适的对应p这个对象的Distance方法来
	   执行。

	   选择器也会被用来选择一个struct类型的字段，比如p.X。由于方法和字段都是在同一
	   命名空间，所以如果我们在这里声明一个X方法的话，编译器会报错，因为在调用p.X时会有
	   歧义(译注：这里确实挺奇怪的)。
	*/
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
	/*
		Path是一个命名的slice类型，而不是Point那样的struct类型，然而我们依然可以为它定义方
		法。在能够给任意类型定义方法这一点上，Go和很多其它的面向对象的语言不太一样。因此
		在Go语言里，我们为一些简单的数值、字符串、slice、map来定义一些附加行为很方便。方
		法可以被声明到任意类型，只要不是一个指针或者一个interface。
	*/
}

/*
基于指针对象的方法
当调用一个函数时，会对其每一个参数值进行拷贝，如果一个函数需要更新一个变量，或者
函数的其中一个参数实在太大我们希望能够避免进行这种默认的拷贝，这种情况下我们就需
要用到指针了。对应到我们这里用来更新接收器的对象的方法，当这个接受者变量本身比较
大时，我们就可以用其指针而不是对象来声明方法
*/
func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

/*
这个方法的名字是 (*Point).ScaleBy  。这里的括号是必须的；没有括号的话这个表达式可能
会被理解为 *(Point.ScaleBy)  。

在现实的程序里，一般会约定如果Point这个类有一个指针作为接收器的方法，那么所有Point
的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数。我们在这里
打破了这个约定只是为了展示一下两种方法的异同而已。
*/

/*
只有类型(Point)和指向他们的指针(*Point)，才是可能会出现在接收器声明里的两种接收器。
此外，为了避免歧义，在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出
现在接收器中的，比如下面这个例子
type P *int
func (P) f() {  } // compile error: invalid receiver type
*/

/*
想要调用指针类型方法 (*Point).ScaleBy  ，只要提供一个Point类型的指针即可，像下面这样。
*/

func main2() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	//
	p := Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	// pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	/*
		不过后面两种方法有些笨拙。幸运的是，go语言本身在这种地方会帮到我们。如果接收器p是
		一个Point类型的变量，并且其方法需要一个Point指针作为接收器，我们可以用下面这种简短
		的写法
	*/
	p.ScaleBy(2)
	/*
	   编译器会隐式地帮我们用&p去调用ScaleBy这个方法。这种简写方法只适用于“变量”，包括
	   struct里的字段比如p.X，以及array和slice内的元素比如perim[0]。我们不能通过一个无法取到
	   地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：
	   Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
	*/
}

/*
下面三种情况里的任意一种情况都是可以的：
不论是接收器的实际参数和其接收器的形式参数相同，比如两者都是类型T或者都是类
型 *T  ：
Point{1, 2}.Distance(q) // Point
pptr.ScaleBy(2) // *Point
或者接收器形参是类型T，但接收器实参是类型 *T  ，这种情况下编译器会隐式地为我们取变
量的地址：
p.ScaleBy(2) // implicit (&p)
或者接收器形参是类型 *T  ，实参是类型T。编译器会隐式地为我们解引用，取到指针指向的
实际变量：
pptr.Distance(q) // implicit (*pptr)
*/

/*******************
其实有两点：
1. 不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型
进行调用的，编译器会帮你做类型转换。
2. 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的内部，第
一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷
贝；第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向
的始终是一块内存地址，就算你对其进行了拷贝。熟悉C或者C艹的人这里应该很快能明
白。
*/
