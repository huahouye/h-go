/*
一个对象的变量或者方法如果对调用方是不可见的话，一般就被定义为“封装”。封装有时候也
被叫做信息隐藏，同时也是面向对象编程最关键的一个方面。
Go语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写
字母的则不会。这种限制包内成员的方式同样适用于struct或者一个类型的方法。因而如果我
们想要封装一个对象，我们必须将其定义为一个struct。
*/
package encapsulation

import ()

/*
封装的第三个优点也是最重要的优点，是阻止了外部调用方对对象内部的值任意地进行修
改。因为对象内部变量只可以被同一个包内的函数修改，所以包的作者可以让这些函数确保
对象内部的一些值的不变性。比如下面的Counter类型允许调用方来增加counter变量的值，并
且允许将这个值reset为0，但是不允许随便设置这个值(译注：因为压根就访问不到)：
*/

type Counter struct{ n int }

func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }

/*
只用来访问或修改内部变量的函数被称为setter或者getter，例子如下，比如log包里的Logger
类型对应的一些函数。在命名一个getter方法时，我们通常会省略掉前面的Get前缀。这种简
洁上的偏好也可以推广到各种类型的前缀比如Fetch，Find或者Lookup。
*/
type Logger struct {
	flags  int
	prefic string
}

func (l *Logger) Flags() int
func (l *Logger) SetFlags(flag int)
func (l *Logger) Prefix() string
func (l *Logger) SetPrefix(prefix string)
