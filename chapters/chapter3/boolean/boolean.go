/*
 &&  的优先级比 ||  高（助记： &&  对应逻辑乘法， ||  对应逻辑加法，乘法比加法优先
级要高）
*/
package main

import ()

func main() {
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		// ...ASCII letter or digit...
	}

	// 布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换
	b := true
	i := 0
	if b {
		i = 1
	}
}

// 如果需要经常做类似的转换, 包装成一个函数会更方便
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// 数字到布尔型的逆转换则非常简单, 不过为了保持对称, 我们也可以包装一个函数
func itob(i int) bool {
	return i != 0
}
