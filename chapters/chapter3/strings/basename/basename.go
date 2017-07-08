package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go"))   // "c.d"
	fmt.Println(basename("abc"))      // "abc"

	fmt.Println(basename2("abc")) // "abc"
}

/*
basename removes directory components and a .suffix.
e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
*/
func basename(s string) string {
	fmt.Println("basename(s string)")
	// Discard last '/' and everything before.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 使用 strings.LastIndex 库函数
func basename2(s string) string {
	fmt.Println("basename2(s string)")
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
