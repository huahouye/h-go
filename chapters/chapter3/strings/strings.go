/*
一个字符串是一个不可改变的字节序列。字符串可以包含任意的数据，包括byte值0，但是通
常是用来包含人类可读的文本。文本字符串通常被解释为采用UTF8编码的Unicode码点
（rune）序列
内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i
个字节的字节值，i必须满足0 ≤ i< len(s)条件约束
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))      // 12
	fmt.Printf("%c\n", s[1]) // e
	fmt.Println(s[0], s[7])  // "104 119" ('h' and 'w')

	// 第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个
	// 字节

	// 子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（并不包含j本身）生成一
	// 个新字符串
	// 左闭右开
	fmt.Println(s[0:5]) // "hello"
	// 同样，如果索引超出字符串范围或者j小于i的话将导致panic异常
	// 不管i还是j都可能被忽略，当它们被忽略时将采用0作为开始位置，采用len(s)作为结束的位置
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	// 其中+操作符将两个字符串链接构造一个新字符串
	fmt.Println("goodbye" + s[5:])

	// 字符串可以用==和<进行比较；比较通过逐个字节比较完成的，因此比较的结果是字符串自然编码的顺序

	// 字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变，当然我们也可以给一
	// 个字符串变量分配一个新字符串值
	ss := "left foot"
	t := ss
	ss += ", right foot"
	// 这并不会导致原始的字符串值被改变，但是变量s将因为+=语句持有一个新的字符串值，但是
	// t依然是包含原先的字符串值
	fmt.Println(s) // "left foot, right foot"
	fmt.Println(t) // "left foot"
	// 因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的
	// s[0] = 'L' // compile error: cannot assign to s[0]

	main2()
	main3()
	main4()
}

/*
不变性意味如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字
符串代价是低廉的。同样，一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地共享
相同的内存，因此字符串切片操作代价也是低廉的。在这两种情况下都没有必要分配新的内
存
*/

/*
Go语言的源文件采用UTF8编码，并且Go语言处理UTF8编码的文本也很出色。unicode包提
供了诸多处理rune字符相关功能的函数（比如区分字母和数组，或者是字母的大写和小写转
换等），unicode/utf8包则提供了用于rune字符序列的UTF8编码和解码的功能。
*/

func main2() {
	fmt.Println("main2()")
	/*
		另一方面，如果我们真的关心每个Unicode字符，我们可以使用其它处理方式。考虑前面的第
		一个例子中的字符串，它包混合了中西两种字符。图3.5展示了它的内存表示形式。字符串包
		含13个字节，以UTF8形式编码，但是只对应9个Unicode字符
	*/
	s := "Hello, 世界"
	fmt.Println(len(s))                    // 13 字节
	fmt.Println(utf8.RuneCountInString(s)) // 9 utf 字符

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// 循环来统计字符串中字符的数目
	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Println(n)
}

func main3() {
	fmt.Println("main3()")
	// string接受到[]rune的类型转换，可以将一个UTF8编码的字符串解码为Unicode字符序列
	// "program" in Japanese katakana
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	// 如果是将一个[]rune类型的Unicode字符slice或数组转为string，则对它们进行UTF8编码
	fmt.Println(string(r)) // "プログラム"

	// 将一个整数转型为字符串意思是生成以只包含对应Unicode码点字符的UTF8字符串
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"

	// 如果对应码点的字符是无效的，则用'\uFFFD'无效字符作为替换
	fmt.Println(string(1234567)) // "?"
}

func main4() {
	fmt.Println("main4()")
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s, b, s2)

	/*
		为了避免转换中不必要的内存分配，bytes包和strings同时提供了许多实用函数。下面是

		strings包中的六个函数
		func Contains(s, substr string) bool
		func Count(s, sep string) int
		func Fields(s string) []string
		func HasPrefix(s, prefix string) bool
		func Index(s, sep string) int
		func Join(a []string, sep string) string

		bytes包中也对应的六个函数
		func Contains(b, subslice []byte) bool
		func Count(s, sep []byte) int
		func Fields(s []byte) [][]byte
		func HasPrefix(s, prefix []byte) bool
		func Index(s, sep []byte) int
		func Join(s [][]byte, sep []byte) []byte
	*/
}
