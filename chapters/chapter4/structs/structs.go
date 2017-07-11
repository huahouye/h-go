/*
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体
*/
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	// dilbert结构体变量的成员可以通过点操作符访问，比如dilbert.Name和dilbert.DoB。因为
	// dilbert是一个变量，它所有的成员也同样是变量，我们可以直接对每个成员赋值
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	// 或者是对成员取地址，然后通过指针访问
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	fmt.Printf("dilbert %v\n", dilbert)

	// 点操作符也可以和指向结构体的指针一起工作
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// 相当于下面语句
	(*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Printf("employeeOfTheMonth %v\n", employeeOfTheMonth)
}

func main1() {
	type Point struct{ x, y int }
	p := Point{1, 2}
	/*
		这里有两种形式的结构体面值语法，上面的是第一种写法，要求以结构体成员定义的顺序为
		每个结构体成员指定一个面值。它要求写代码和读代码的人要记住结构体的每个成员的类型
		和顺序，不过结构体成员有细微的调整就可能导致上述代码不能编译。因此，上述的语法一
		般只在定义结构体的包内部使用，或者是在较小的结构体中使用，这些结构体的成员排列比
		较规则，比如image.Point{x, y}或color.RGBA{red, green, blue, alpha}。
		其实更常用的是第二种写法，以成员名字和相应的值来初始化，可以包含部分或全部的成
		员，如1.4节的Lissajous程序的写法
	*/
	// anim := gif.GIF{LoopCount: nframes}
}
