/*
一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相
同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以
继续执行后面的语句。反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有
另一个goroutine在相同的Channels上执行发送操作。

基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。因为这个原
因，无缓存Channels有时候也被称为同步Channels。当通过一个无缓存Channels发送数据
时，接收者收到数据发生在唤醒发送者goroutine之前（译注：happens before，这是Go语言
并发内存模型的一个关键术语！）。

在讨论并发编程时，当我们说x事件在y事件之前发生（happens before），我们并不是说x事
件在时间上比y时间更早；我们要表达的意思是要保证在此之前的事件都已经完成了，例如在
此之前的更新某些变量的操作已经完成，你可以放心依赖这些已完成的事件了。
当我们说x事件既不是在y事件之前发生也不是在y事件之后发生，我们就说x事件和y事件是并
发的。这并不是意味着x事件和y事件就一定是同时发生的，我们只是不能确定这两个事件发
生的先后顺序。在下一章中我们将看到，当两个goroutine并发访问了相同的变量时，我们有
必要保证某些事件的执行顺序，以避免出现某些并发问题。

在8.3节的客户端程序，它在主goroutine中（译注：就是执行main函数的goroutine）将标准输
入复制到server，因此当客户端程序关闭标准输入时，后台goroutine可能依然在工作。我们
需要让主goroutine等待后台goroutine完成工作后再退出，我们使用了一个channel来同步两个
goroutine
*/
package main

import (
	"fmt"
)

func main() {
	// 以最简单方式调用make函数创建的时一个无缓存的channel
	ch := make(chan int) // ch has type 'chan int'
	// 我们也可以指定第二个整
	// 形参数，对应channel的容量。如果channel的容量大于零，那么该
	// channel 就是带缓存的 channel。
	ch1 = make(chan int)    // unbuffered channel
	ch2 = make(chan int, 0) // unbuffered channel
	ch3 = make(chan int, 3) // buffered channel with capacity 3

	/*
		一个channel有发送和接受两个主要操作，都是通信行为。一个发送语句将一个值从一个
		goroutine通过channel发送到另一个执行接收操作的goroutine。发送和接收两个操作都是
		用 <-  运算符。在发送语句中， <-  运算符分割channel和要发送的值。在接收语句中， <-  运
		算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。
	*/
	x := 3
	ch <- x  // a send statement
	x = <-ch // a receive expression in an assignment statement
	<-ch     // a receive statement; result is discarded
	close(ch)
}

/*
基于channels发送消息有两个重要方面。首先每个消息都有一个值，但是有时候通讯的事实
和发生的时刻也同样重要。当我们更希望强调通讯发生的时刻时，我们将它称为消息事件。
有些消息事件并不携带额外的信息，它仅仅是用作两个goroutine之间的同步，这时候我们可
以用 struct{}  空结构体作为channels元素的类型，虽然也可以使用bool或int类型实现同样的
功能， done <- 1  语句也比 done <- struct{}{}  更短。
*/
func main2() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}
