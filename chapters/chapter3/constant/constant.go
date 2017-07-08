package main

import (
	"fmt"
)

func main() {
	const pi = 3.14159 // approximately; math.Pi is a better approximatio
	// 和变量声明一样，可以批量声明多个常量；这比较适合声明一组相关的常量
	const (
		e   = 2.71828182845904523536028747135266249775724709369995957496696763
		pi2 = 3.14159265358979323846264338327950288419716939937510582097494459
	)

	fmt.Println(pi, e, pi2)

	main2()
}

func main1() {
	// 义了一个Weekday命名类型，然后为一周的每天定义了
	// 一个常量，从周日0开始。在其它编程语言中，这种类型一般被称为枚举类型
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
}

/*
我们也可以在复杂的常量表达式中使用iota，下面是来自net包的例子，用于给一个无符号整
数的最低5bit的每个bit指定一个名字
*/
type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

/*
		随着iota的递增，每个常量对应表达式1 << iota，是连续的2的幂，分别对应一个bit位置。使
	用这些常量可以用于测试、设置或清除对应的bit位的值
*/
func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main2() {
	fmt.Println("main2()")
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

////////////////
// 下面是一个更复杂的例子，每个常量都是1024的幂
// 不过iota常量生成规则也有其局限性。例如，它并不能用于产生1000的幂（KB、MB等），因
// 为Go语言并没有计算幂的运算符。
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)
