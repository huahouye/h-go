/*
左移运算用零填充右边空缺的bit位，无符号数的右移运算也是用0填充左边空缺的bit位，但是
有符号数的右移运算会用符号位的值填充左边空缺的bit位。因为这个原因，最好用无符号运
算，这样你可以将整数完全当作一个bit位模式处理。
*/
package main

import (
	"fmt"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	// 检测 bit 位为 1 的位置
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	main2()
	main3()
	main4()
	main5()
}

func main2() {
	/*
			尽管Go语言提供了无符号数和运算，即使数值本身不可能出现负数我们还是倾向于使用有符
		号的int类型，就像数组的长度那样，虽然使用uint无符号类型似乎是一个更合理的选择。事实
		上，内置的len函数返回一个有符号的int，我们可以像下面例子那样处理逆序循环。
	*/
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
	/*
			另一个选择对于上面的例子来说将是灾难性的。如果len函数返回一个无符号数，那么i也将是
		无符号的uint类型，然后条件 i >= 0  则永远为真。在三次迭代之后，也就是 i == 0  时，i--语
		句将不会产生-1，而是变成一个uint类型的最大值（可能是2 4 − 1），然后medals[i]表达式将
		发生运行时panic异常（§5.9），也就是试图访问一个slice范围以外的元素。
	*/
}

func main3() {
	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // 3.141 3

	f = 1.99
	fmt.Println(int(f)) // 1
}

func main4() {
	// 任何大小的整数字面值都可以用以0开始的八进制格式书写，例如0666
	// 或用以0x或0X开头的十六进制格式书写，例如0xdeadbeef。十六进制数字可以用大写或小写字母
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	/*
			请注意 fmt 的两个使用技巧。通常 Printf 格式化字符串包含多个 % 参数时将会包含对应相同数量
		的额外操作数，但是%之后的 [1] 副词告诉 Printf 函数再次使用第一个操作数。第二，%后
		的 # 副词告诉 Printf 在用 %o、%x 或 %X 输出时生成 0、0x 或 0X 前缀。
	*/
}

func main5() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}
