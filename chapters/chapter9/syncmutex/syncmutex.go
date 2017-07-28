/*
一个buffered channel作为一个计数信号量，来保证最多只有20个
goroutine会同时执行HTTP请求。同理，我们可以用一个容量只有1的channel来保证最多只有
一个goroutine在同一时刻访问一个共享变量。一个只能为1和0的信号量叫做二元信号量
*/
package syncmutex

import (
	"sync"
)

// 二元信号量
var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}
func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

// Lock
var (
	mu      sync.Mutex // guards balance
	balance int
)

// goroutine在结束后释放锁是必要的，无论以哪条路径通过函数都需要释放，即使是在
// 错误路径中，也要记得释放。
func Deposit(amount int) {
	mu.Lock()
	// 惯例来说，被mutex所保护的变量是在mutex变量声明之后立刻声明的
	balance = balance + amount
	mu.Unlock()
	// 在Lock和Unlock之间的代码段中的内容goroutine可以随便读取或者修改，这个代码段叫做临
	// 界区。
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

/*
上面的bank程序例证了一种通用的并发模式。一系列的导出函数封装了一个或多个变量，那
么访问这些变量唯一的方式就是通过这些函数来做(或者方法，对于一个对象的变量来说)。每
一个函数在一开始就获取互斥锁并在最后释放锁，从而保证共享变量不会被并发访问。这种
函数、互斥锁和变量的编排叫作监控monitor(这种老式单词的monitor是受"monitor
goroutine"的术语启发而来的。两种用法都是一个代理人保证变量被顺序访问)。
*/

/*
我们用defer来调用Unlock，临界区会
隐式地延伸到函数作用域的最后，这样我们就从“总要记得在函数返回之后或者发生错误返回
时要记得调用一次Unlock”这种状态中获得了解放。Go会自动帮我们完成这些事情。
*/
func Balance2() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
