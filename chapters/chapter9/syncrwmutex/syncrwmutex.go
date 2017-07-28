/*
“多读单写”锁(multiple readers, single writer lock),sync.RWMutex

RWMutex只有当获得锁的大部分goroutine都是读操作，而锁在竞争条件下，也就是说，
goroutine们必须等待才能获取到锁的时候，RWMutex才是最能带来好处的。RWMutex需要更
复杂的内部记录，所以会让它比一般的无竞争锁的mutex慢一些。

尽管去理解并发的一种尝试是去将其运行理解为不同goroutine语句的交错执行，但看看上面
的例子，这已经不是现代的编译器和cpu的工作方式了。因为赋值和打印指向不同的变量，编
译器可能会断定两条语句的顺序不会影响执行结果，并且会交换两个语句的执行顺序。如果
两个goroutine在不同的CPU上执行，每一个核心有自己的缓存，这样一个goroutine的写入对
于其它goroutine的Print，在主存同步之前就是不可见的了。
所有并发的问题都可以用一致的、简单的既定的模式来规避。所以可能的话，将变量限定在
goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问。
*/
package syncrwmutex

import (
	"sync"
)

var mu sync.RWMutex
var balance int

// RLock 只能在临界区共享变量没有任何写入操作时可用。
func Balance() int {
	mu.RLock() // readers lock
	defer mu.RUnlock()
	return balance
}
