/*
一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进
制数的精度；通常应该优先使用float64类型，因为float32类型的累计计算误差很容易扩散，
并且float32能精确表示的正整数并不是很大（译注：因为float32的有效bit位只有23个，其它
的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差）
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	main1()
	main2()
	main3()
	v1, b1 := main4(true)
	fmt.Println(v1, b1)
	v2, b2 := main4(false)
	fmt.Println(v2, b2)
}

func main1() {
	fmt.Println("main1()")
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}

func main2() {
	fmt.Println("main2()")
	/*
			用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是
		对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适。所有的这三个打印形式都
		可以指定打印的宽度和控制打印精度。
	*/
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func main3() {
	fmt.Println("main3()")
	/*
			math包中除了提供大量常用的数学函数外，还提供了IEEE754浮点数标准中定义的特殊值的
		创建和测试：正无穷大和负无穷大，分别用于表示太大溢出的数字和除零的结果；还有NaN
		非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1).
	*/
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
}

/*
如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败，像这样
*/
func main4(flag bool) (value float64, ok bool) {
	fmt.Println("main4()")
	if !flag {
		return 0, false
	}
	return 1.0, true
}
