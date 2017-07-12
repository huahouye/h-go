/*
1、内置的error是接口类型。我们将在第七章了解接口类型的含义，以及它对错误处理的影响。
现在我们只需要明白error类型可能是nil或者non-nil。nil意味着函数运行成功，non-nil表示失
败。对于non-nil的error类型,我们可以通过调用error的Error函数或者输出函数获得字符串类型
的错误信息。

2、在Go中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常
（exception），这使得Go有别于那些将函数运行失败看作是异常的语言。虽然Go有各种异
常机制，但这些机制仅被使用在处理那些未被预料到的错误，即bug，而不是那些在健壮程序
中应该被避免的程序错误。

3、错误处理策略
3.1、首先，也是最常用的方式是传播错误。这意味着函数中某个子程序的失败，会变成该函数的失败。
fmt.Errorf函数使用fmt.Sprintf格式化错误信息并返回。我们使用该函数前缀添加额外的上下文
信息到原始错误信息。当错误最终由main函数处理时，错误信息应提供清晰的从原因到后果
的因果链，就像美国宇航局事故调查时做的那样：
genesis: crashed: no parachute: G-switch failed: bad relay orientation
由于错误信息经常是以链式组合在一起的，所以错误信息中应避免大写和换行符。
3.2、第二种策略。如果错误的发生是偶然性的，或由不可预知的问题导
致的。一个明智的选择是重新尝试失败的操作。在重试时，我们需要限制重试的时间间隔或
重试的次数，防止无限制的重试。
3.3、如果错误发生后，程序无法继续运行，我们就可以采用第三种策略：输出错误信息并结束程
序。需要注意的是，这种策略只应在main中执行。对库函数而言，应仅向上传播错误，除非
该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。
3.4、第四种策略：有时，我们只需要输出错误信息就足够了，不需要中断程序的运行。我们可以
通过log包提供函数
log.Printf("ping failed: %v; networking disabled",err)
3.5、第五种，也是最后一种策略：我们可以直接忽略掉错误。
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	err()
	WaitForServer("http://www.baidu.com")
}

func err() {
	fmt.Println(err)
	fmt.Printf("%v", err)
}

// 3.2、第二种策略
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	fmt.Println("WaitForServer(url string)")
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
