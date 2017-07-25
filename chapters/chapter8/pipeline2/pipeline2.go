/*
Go语言的类型系统提供了单方向的channel类型，分别用于
只发送或只接收的channel。类型 chan<- int  表示一个只发送int的channel，只能发送不能接
收。相反，类型 <-chan int  表示一个只接收int的channel，只能接收不能发送。（箭头 <-  和
关键字chan的相对位置表明了channel的方向。）这种限制将在编译期检测。
-- <- 所在位置 左收右发
*/
package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
