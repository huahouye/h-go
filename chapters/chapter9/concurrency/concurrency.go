/*
有三种方式可以避免数据竞争：
第一种方法是不要去写变量。如果我们在创建goroutine之前的初始化阶段，就初始化了map中的所有条目并且再也
不去修改它们，那么任意数量的goroutine并发访问Icon都是安全的，因为每一个goroutine都
只是去读取而已。
第二种避免数据竞争的方法是，避免从多个goroutine访问变量。由于其它的goroutine不能够直接访问变量，
它们只能使用一个channel来发送给指定的
goroutine请求来查询更新变量。这也就是Go的口头禅“不要使用共享数据来通信；使用通信来
共享数据”。一个提供对一个指定的变量通过cahnnel来请求的goroutine叫做这个变量的监控
(monitor)goroutine。
第三种避免数据竞争的方法是允许很多goroutine去访问变量，但是在同一个时刻最多只有一
个goroutine在访问。这种方式被称为“互斥”
*/
package concurrency

import ()

// 方法一
var icons = map[string]image.Image{
	"spades.png":   loadIcon("spades.png"),
	"hearts.png":   loadIcon("hearts.png"),
	"diamonds.png": loadIcon("diamonds.png"),
	"clubs.png":    loadIcon("clubs.png"),
}

// Concurrency-safe.
func Icon(name string) image.Image { return icons[name] }

// 方法二
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
func Deposit(amount int)      { deposits <- amount }
func Balance() int            { return <-balances }
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
func init() {
	go teller() // start the monitor goroutine
}

// 方法三
type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}
func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touches this cake again
	}
}
