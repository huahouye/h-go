/*
nonempty函数将在原有slice内存空间之上返回不包含空字符串的列表
Nonempty is an example of an in-place slice algorithm.
*/
package main

import (
	"fmt"
)

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func main() {
	data := []string{"one", "", "three"}
	// 比较微妙的地方是，输入的slice和输出的slice共享一个底层数组。这可以避免分配另一个数
	// 组，不过原来的数据将可能会被覆盖，正如下面两个打印语句看到的那样
	fmt.Printf("%q\n", noneempty(data)) // ["one" "three"]
	fmt.Printf("%q\n", data)            // ["one" "three" "three"]
}

func noneempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
