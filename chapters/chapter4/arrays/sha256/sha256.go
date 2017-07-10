package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x")) // 小写 x
	c2 := sha256.Sum256([]byte("X")) // 大写 X
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	// Printf函数的%x副词参数，它用于指定以十六进制的格式打印数组或
	// slice全部的元素，%t副词参数是用于打印布尔型数据，%T副词参数是用于显示一个值对应的
	//数据类型

	// 传入一个数组指针
	ba := [32]byte{} // 其实这样就可以生成一个32字节的数组
	zero(&ba)
}

func zero(ptr *[32]byte) {
	fmt.Println("zero(ptr *[32]byte)")
	/*
		我们可以显式地传入一个数组指针，那样的话函数通过指针对数组的任何修改都可以
		直接反馈到调用者。下面的函数用于给[32]byte类型的数组清零
	*/
	for i := range ptr { // 或者 *ptr = [32]byte{}
		ptr[i] = 0
	}
	fmt.Println(*ptr)
}

/*
像SHA256这类需要处理特定大小数组的特例外，数组依然很少用作函数参数；相反，
我们一使用slice来替代数组。
*/
