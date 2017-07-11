/*
一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。
（该限制同样适应于数组。）但是S类型的结构体可以包含 *S  指针类型的成员，这可以让我
们创建递归的数据结构，比如链表和树结构等。在下面的代码中，我们使用一个二叉树来实
现一个插入排序
*/
package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {

}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

/*
seen := make(map[string]struct{}) // set of strings
// ...
if _, ok := seen[s]; !ok {
seen[s] = struct{}{}
// ...first time seeing s...
}
*/
