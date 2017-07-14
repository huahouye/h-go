package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
db.list的调用会援引一个接收者是db的database.list方法。所以db.list是一个实现了
handler类似行为的函数，但是因为它没有方法，所以它不满足http.Handler接口并且不能直接
传给mux.Handle。
*/
func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	fmt.Println("server listen and serve on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

/*
http://localhost:8000/list
http://localhost:8000/price?item=socks
*/
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

/*
语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，因为http.HandlerFunc是一个类
型
package http
type HandlerFunc func(w ResponseWriter, r *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
*/

/*
HandlerFunc显示了在Go语言接口机制中一些不同寻常的特点。这是一个有实现了接口
http.Handler方法的函数类型。ServeHTTP方法的行为调用了它本身的函数。因此
HandlerFunc是一个让函数值满足一个接口的适配器，这里函数和这个接口仅有的方法有相同
的函数签名。实际上，这个技巧让一个单一的类型例如database以多种方式满足http.Handler
接口：一种通过它的list方法，一种通过它的price方法等等。
因为handler通过这种方式注册非常普遍，ServeMux有一个方便的HandleFunc方法，它帮我
们简化handler注册代码成这样
mux.HandleFunc("/list", db.list)
mux.HandleFunc("/price", db.price)
*/

/*
从上面的代码很容易看出应该怎么构建一个程序，它有两个不同的web服务器监听不同的端口
的，并且定义不同的URL将它们指派到不同的handler。我们只要构建另外一个ServeMux并且
在调用一次ListenAndServe（可能并行的）。但是在大多数程序中，一个web服务器就足够
了。此外，在一个应用程序的多个文件中定义HTTP handler也是非常典型的，如果它们必须
全部都显示的注册到这个应用的ServeMux实例上会比较麻烦。
所以为了方便，net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的
http.Handle和http.HandleFunc函数。现在，为了使用DefaultServeMux作为服务器的主
handler，我们不需要将它传给ListenAndServe函数；nil值就可以工作。
然后服务器的主函数可以简化成
*/

func main1() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
