package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

}

func title1(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Check Content-Type is HTML (e.g., "text/html;charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

/*
你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。当defer语
句被执行时，跟在defer后面的函数会被延迟执行。直到包含该defer语句的函数执行完毕时，
defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic
导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相
反。
defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通
过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。释放资源的
defer应该直接跟在请求资源的语句后。在下面的代码中，一条defer语句替代了之前的所有
resp.Body.Close
*/
func title2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// 直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，不论包
	// 含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// ...print doc's title element…
	return nil
}

// 在处理其他资源时，也可以采用defer机制，比如对文件的操作
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadAll(f)
}
func ReadAll(file *os.File) ([]byte, error) {
	return nil, nil
}

// 处理互斥锁
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

/*
调试复杂程序时，defer机制也常被用于记录何时进入和退出函数。下例中的
bigSlowOperation函数，直接调用trace记录函数的被调情况。bigSlowOperation被调时，
trace会返回一个函数值，该函数值会在bigSlowOperation退出时被调用。通过这种方式， 我
们可以只通过一条语句控制函数的入口和所有的出口，甚至可以记录函数的运行时间，如例
子中的start。需要注意一点：不要忘记defer语句后的圆括号，否则本该在进入时执行的操作
会在退出时执行，而本该在退出时执行的，永远不会被执行。

defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义
的匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，对匿名函数采用defer机
制，可以使其观察函数的返回值
*/
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
