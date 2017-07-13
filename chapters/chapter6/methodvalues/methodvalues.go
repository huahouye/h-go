/*
我们经常选择一个方法，并且在同一个表达式里执行，比如常见的p.Distance()形式，实际上
将其分成两步来执行也是可能的。p.Distance叫作“选择器”，选择器会返回一个方法"值"->一
个将方法(Point.Distance)绑定到特定接收器变量的函数。这个函数可以不通过指定其接收器
即可被调用；即调用时不需要指定接收器(译注：因为已经在前文中指定过了)，只要传入函数
的参数即可
*/
package methodvalues

import (
	"fmt"
	"time"
)

type Point struct {
	X, Y float64
}

func main1() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)
	scaleP := p.ScaleBy                // method value
	scaleP(2)                          // p becomes (2, 4)
	scaleP(3)                          // then (6, 12)
	scaleP(10)                         // then (60, 120)
}

type Rocket struct {
}

func (r *Rocket) Launch() {

}
func main2() {
	r := new(Rocket)
	time.AfterFunc(10*time.Second, func() { r.Launch() })

	// 直接用方法"值"传入AfterFunc的话可以更为简短
	time.AfterFunc(10*time.Second, r.Launch)
}
