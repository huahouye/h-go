/*
哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key
都是不同的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。

在Go语言中，一个map就是一个哈希表的引用，map类型可以写为map[K]V，其中K和V分别
对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，但是
key和value之间可以是不同的数据类型。其中K对应的key必须是支持==比较运算符的数据类
型，所以map可以通过测试key是否相等来判断是否已经存在。虽然浮点数类型也是支持相等
运算符比较的，但是将浮点数用做key类型则是一个坏的想法，正如第三章提到的，最坏的情
况是可能出现的NaN和任何浮点数都不相等。对于V对应的value数据类型则没有任何的限
制。
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int) // mapping from strings to ints
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// 相当于
	ages3 := make(map[string]int)
	ages3["alice"] = 31
	ages3["charlie"] = 34
	fmt.Println(ages, ages2, ages3)

	delete(ages3, "alice") // remove element ages["alice"]
	fmt.Println(ages, ages2, ages3)

	// 所有这些操作是安全的，即使这些元素不在map中也没有关系；如果一个查找失败将返回
	// value类型对应的零值，例如，即使map中不存在“bob”下面的代码也可以正常工作，因为
	// ages["bob"]失败时将返回0。
	ages["bob"] = ages["bob"] + 1 // happy birthday!
	// 而且 x += y  和 x++  等简短赋值语法也可以用在map上，所以上面的代码可以改写成
	ages["bob"] += 1
	// 更简单的写法
	ages["bob"]++
	fmt.Println(ages["bob"])

	// 但是map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：
	// _ = &ages["bob"] // compile error: cannot take address of map element
	// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而
	// 可能导致之前的地址无效。

	/*
		Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。在实践
		中，遍历的顺序是随机的，每一次遍历的顺序都不相同。这是故意的，每次都使用随机的遍
		历顺序可以强制要求程序不会依赖具体的哈希函数实现。如果要按顺序遍历key/value对，我
		们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。下面是
		常见的处理方式
	*/
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	main1()
	// True if equal is written incorrectly.
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

func main1() {
	fmt.Println("main1()")
	var names []string
	fmt.Println(names == nil)    // "true"
	fmt.Println(len(names) == 0) // "true"
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	/*
		map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它
		们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常
		ages["carol"] = 21 // panic: assignment to entry in nil map
		在向map存数据前必须先创建map。
	*/

	/*
		有时候可能需要知道对应的元素是否真的是在map
		之中。例如，如果元素类型是一个数字，你可以需要区分一个已经存在的0，和不存在而返回
		零值的0，可以像下面这样测试
		age, ok := ages["bob"]
		if !ok {} // "bob" is not a key in this map; age == 0.
		// 你会经常看到将这两个结合起来使用，像这样
		if age, ok := ages["bob"]; !ok {}

		map的下标语法将产生两个值；第二个是一个布尔值，用于报告元素是否真
		的存在。布尔变量一般命名为ok，特别适合马上用于if条件判断部分。
	*/
}

/*
和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。要判断两个map是
否包含相同的key和value，我们必须通过一个循环实现：
*/
func equal(x, y map[string]int) bool {
	fmt.Println("equal(x, y map[string]int)")
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
