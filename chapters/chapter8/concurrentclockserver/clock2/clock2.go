package main

import (
	"io"
	"log"
	"net"
	"time"
)

// nc localhost 8000
// telnet localhost 8000
func main() {
	/*
		Listen函数创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接，在这
		个例子里我们用的是TCP的localhost:8000端口。listener对象的Accept方法会直接阻塞，直到
		一个新的连接被创建，然后会返回一个net.Conn对象来表示这个连接。
	*/
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}
}

/*
handleConn函数会处理一个完整的客户端连接。在一个for死循环中，将当前的时候用
time.Now()函数得到，然后写到客户端。由于net.Conn实现了io.Writer接口，我们可以直接向
其写入内容。这个死循环会一直执行，直到写入失败。最可能的原因是客户端主动断开连
接。这种情况下handleConn函数会用defer调用关闭服务器侧的连接，然后返回到主函数，继
续等待下一个连接请求。
*/
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		/*
			time.Time.Format方法提供了一种格式化日期和时间信息的方式。它的参数是一个格式化模板
			标识如何来格式化时间，而这个格式化模板限定为Mon Jan 2 03:04:05PM 2006 UTC-0700。
			有8个部分(周几，月份，一个月的第几天，等等)。可以以任意的形式来组合前面这个模板；
			出现在模板中的部分会作为参考来对时间格式进行输出。
		*/
		/*
			这是go语言和其它语言相比比较奇葩的
			一个地方。。你需要记住格式化字符串是1月2日下午3点4分5秒零六年UTC-0700，而不像其
			它语言那样Y-m-d H:i:s一样，当然了这里可以用1234567的方式来记忆，倒是也不麻烦
		*/
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
