// 需要注意，Go语言并不需要显式地在每一个case后写break，语言默认执行完
// case后的逻辑语句会自动退出。当然了，如果你想要相邻的几个case都执行同一逻辑的话，
// 需要自己显式地写上一个fallthrough语句来覆盖这种默认行为。不过fallthrough语句在一般的
// 程序中很少用到。
package main

import (
	"fmt"
)

func main() {
	var heads, tails int
	switch coinflip() {
	case "heads":
		heads++
		fmt.Printf("heads %d\n", heads)
	case "tails":
		tails++
		fmt.Printf("heads %d\n", tails)
	default:
		fmt.Println("landed on edge!")
	}
}

func coinflip() string {
	return "heads"
}

// Go语言里的switch还可以不带操作对象（译注：switch不带操作对象时默认用true值代替，然
// 后将每个case的表达式和true值进行比较）；可以直接罗列多种条件，像其它语言里面的多个
// if else一样
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
