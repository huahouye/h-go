/*
Channels也可以用于将多个goroutine链接在一起，一个Channels的输出作为下
一个Channels的输入。这种串联的Channels就是所谓的管道（pipeline）
*/
package main

import (
	"fmt"
)

func main() {
	main2()
	main1()
}

func main1() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 5; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(naturals)
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

/*
因为上面的语法是笨拙的，而且这种处理模式很场景，因此Go语言的range循环
可直接在 channels 上面迭代。
*/
func main2() {
	fmt.Println("main1()")

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
