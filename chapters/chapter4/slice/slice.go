/*
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作
[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。
数组和slice之间有着紧密的联系。一个slice是一个轻量级的数据结构，提供了访问数组子序
列（或者全部）元素的功能，而且slice的底层确实引用一个数组对象。一个slice由三个部分
构成：
指针、长度和容量。
指针指向第一个slice元素对应的底层数组元素的地址，要注意的
是slice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目；长度不
能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分
别返回slice的长度和容量。

slice的切片操作s[i:j]，其中0 ≤ i≤ j≤ cap(s)，用于创建一个新的slice，引用s的从第i个元素开
始到第j-1个元素的子序列。新的slice将只有j-i个元素。如果i位置的索引被省略的话将使用0代
替，如果j位置的索引被省略的话将使用len(s)代替。
*/
package main

import (
	"fmt"
)

// 要注意的是slice类型的变量s和数组类型的变量a的初始化语法的差异
// slice和数组的字面值
// 语法很类似，它们都是用花括弧包含一系列的初始化元素，但是对于slice并没有指明序列的
// 长度。这会隐式地创建一个合适大小的数组，然后slice的指针指向底层的数组。就像数组字
// 面值一样，slice的字面值也可以按顺序指定初始化值序列，或者是通过索引和元素值指定，
// 或者的两种风格的混合语法初始化。
func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	// 一种将slice元素循环向左旋转n个元素的方法是三次调用reverse反转函数，第一次是反转开头
	// 的n个元素，然后是反转剩下的元素，最后是反转整个slice的元素。（如果是向右循环旋转，
	// 则将第三个函数调用移到第一个调用位置就可以了。）
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	sa1 := []string{"xxx", "yyy", "zzz"}
	sa2 := []string{"xxx", "yyy", "zzz"}
	rs := equal(sa1, sa2)
	fmt.Println(rs)

	compareNil()
	appendFunc()
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
和数组不同的是，slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有
全部相等元素。不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相
等（[]byte），但是对于其他类型的slice，我们必须自己展开每个元素进行比较：
*/
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
	/*
		上面关于两个slice的深度相等测试，运行的时间并不比支持==操作的数组或字符串更多，但
		是为何slice不直接支持比较运算符呢？这方面有两个原因。第一个原因，一个slice的元素是
		间接引用的，一个slice甚至可以包含自身。虽然有很多办法处理这种情形，但是没有一个是
		简单有效的。
		第二个原因，因为slice的元素是间接引用的，一个固定值的slice在不同的时间可能包含不同
		的元素，因为底层数组的元素可能会被修改。并且Go语言中map等哈希表之类的数据结构的
		key只做简单的浅拷贝，它要求在整个声明周期中相等的key必须对相同的元素。对于像指针
		或chan之类的引用类型，==相等测试可以判断两个是否是引用相同的对象。一个针对slice的
		浅相等测试的==操作符可能是有一定用处的，也能临时解决map类型的key问题，但是slice和
		数组不同的相等测试行为会让人困惑。因此，安全的做法是直接禁止slice之间的比较操作。
	*/

}

func compareNil() {
	fmt.Println("compareNil()")
	// slice唯一合法的比较操作是和nil比较，例如
	// if summer == nil { /* ... */ }
	// 一个零值的slice等于nil。一个nil值的slice并没有底层数组。一个nil值的slice的长度和容量都
	// 是0，但是也有非nil值的slice的长度和容量也是0的，例如[]int{}或make([]int, 3)[3:]。与任意类
	// 型的nil值一样，我们可以用[]int(nil)类型转换表达式来生成一个对应类型slice的nil值。
	var s []int
	fmt.Println(s == nil, len(s)) // len(s) == 0, s == nil
	s = nil
	fmt.Println(s == nil, len(s)) // len(s) == 0, s == nil
	s = []int(nil)
	fmt.Println(s == nil, len(s)) // len(s) == 0, s == nil
	s = []int{}
	fmt.Println(s == nil, len(s)) // len(s) == 0, s != nil

	/*
	   如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。
	   除了和nil相等比较外，一个nil值的slice的行为和其它任意0长度的slice一样；例如reverse(nil)
	   也是安全的。除了文档已经明确说明的地方，所有的Go语言函数应该以相同的方式对待nil值
	   的slice和0长度的slice。
	*/

	// 内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情
	// 况下，容量将等于长度。
	ia3 := make([]int, 3)
	iac := make([]int, 3, 3) // same as make([]T, cap)[:len]
	fmt.Println(ia3, iac)
	// 在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引
	// 用底层匿名的数组变量。在第一种语句中，slice是整个数组的view。在第二个语句中，slice
	// 只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增
	// 长用的。
}

// 内置的append函数用于向slice追加元素
func appendFunc() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	// 可以通过Go语言内置的[]rune("Hello, 世界")转换操作完成
	fmt.Printf("%q\n", runes) // // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}
