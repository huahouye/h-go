// 开头家注释，对整个文件的源码做描述
package main

import (
	"fmt"
	"os"
)

func main() {
	// version 1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// version 2
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s1 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)

	// 高效的字符串拼接
	fmt.Println(strings.Join(os.Args[1:], " "))
}
