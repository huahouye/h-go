/*
如果初始化成本比较大的话，那么将初始化延迟到需要的时候再去做就是一个比较好的选择。
*/
package synconce

import (
	"image"
	"sync"
)

var icons map[string]image.Image

// 1、懒初始化
func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // one-time initialization
	}
	return icons[name]
}

// 2、最简单且正确的保证所有goroutine能够观察到loadIcons效果的方式，是用一个mutex来同步检查。
var mu sync.Mutex

// Concurrency-safe.
func Icon(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

// 3、这里我们可以引入一个允许多读的锁
var rwmu = sync.RWMutex

// Concurrency-safe.
func Icon3(name string) image.Image {
	rwmu.RLock()
	if icons != nil {
		icon := icons[name]
		rwmu.RUnlock()
		return icon
	}
	rwmu.RUnlock()

	// acquire an exclusive lock
	rwmu.Lock()
	if icons == nil { // NOTE: must recheck for nil
		loadIcons()
	}
	icon := icons[name]
	rwmu.Unlock()
	return icon
}

/*
上面的代码有两个临界区。goroutine首先会获取一个写锁，查询map，然后释放锁。如果条
目被找到了(一般情况下)，那么会直接返回。如果没有找到，那goroutine会获取一个写锁。不
释放共享锁的话，也没有任何办法来将一个共享锁升级为一个互斥锁，所以我们必须重新检
查icons变量是否为nil，以防止在执行这一段代码的时候，icons变量已经被其它gorouine初始
化过了。
*/

// 4.sync.Once
/*
一次性初始化的问题：sync.Once。概念上来讲，
一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成了；
互斥量用来保护boolean变量和客户端数据结构。Do这个唯一的方法需要接收初始化函数作为
其参数。
*/
var loadIconsOnce sync.Once
var icons map[string]image.Image

// Concurrency-safe.
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

/*
每一次对Do(loadIcons)的调用都会锁定mutex，并会检查boolean变量。在第一次调用时，变
量的值是false，Do会调用loadIcons并会将boolean设置为true。随后的调用什么都不会做，
但是mutex同步会保证loadIcons对内存(这里其实就是指icons变量啦)产生的效果能够对所有
goroutine可见。用这种方式来使用sync.Once的话，我们能够避免在变量被构建完成之前和其
它goroutine共享该变量。
*/
