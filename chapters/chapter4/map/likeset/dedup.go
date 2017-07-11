/*
Go 语言中并没有提供一个 set 类型，但是 map 中的 key 也是不相同的，可以用map实现类似set
的功能。为了说明这一点，下面的dedup程序读取多行输入，但是只打印第一次出现的行。
（它是 1.3 节中出现的 dup 程序的变体。）dedup程序通过map来表示所有的输入行所对应的
set集合，以确保已经在集合存在的行不会被重复打印。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	}
}
