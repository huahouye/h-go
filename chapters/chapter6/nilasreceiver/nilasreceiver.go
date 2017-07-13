/*
就像一些函数允许nil指针作为参数一样，方法理论上也可以用nil指针作为其接收器，尤其当
nil对于对象来说是合法的零值时，比如map或者slice。在下面的简单int链表的例子里，nil代
表的是空链表：
*/
package nilasreceiver

import ()

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
